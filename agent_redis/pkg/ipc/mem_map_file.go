package ipc

import (
	"io"
	"os"
	"runtime"
	"sync"
	"syscall"
)

type IMemMapFile interface {
	io.WriteCloser
}

type memMapFileImpl struct {
	ptr []byte
	size int64
	mutex *sync.Mutex
}

func (mmfi *memMapFileImpl)Write(b []byte) (n int, err error) {
	if mmfi.ptr == nil || mmfi.mutex == nil {
		err = os.ErrInvalid
		return
	}
	mmfi.mutex.Lock()
	defer mmfi.mutex.Unlock()

	n = copy(mmfi.ptr, b)
	return
}

func (mmfi *memMapFileImpl)Close() error {
	m := mmfi.mutex
	m.Lock()
	data := mmfi.ptr
	mmfi.mutex = nil
	mmfi.ptr = nil
	m.Unlock()

	runtime.SetFinalizer(mmfi, nil)
	return syscall.Munmap(data)
}

func MemMapFileOpen(filename string, size int64) (IMemMapFile, error) {
	f,openErr := os.Open(filename)
	if openErr != nil {
		return nil, openErr
	}
	defer f.Close()

	resizeErr := f.Truncate(size)
	if resizeErr != nil {
		return nil, resizeErr
	}

	data, mmapErr := syscall.Mmap(int(f.Fd()), 0, int(size), syscall.PROT_WRITE, syscall.MAP_SHARED);
	if mmapErr != nil {
		return nil, mmapErr
	}
	ret := &memMapFileImpl{
		ptr : data,
		size : size,
		mutex: new(sync.Mutex),
	}
	runtime.SetFinalizer(ret, ret.Close())
	return ret,nil
}