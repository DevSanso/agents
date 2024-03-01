package ipc

import (
	"agent_redis/pkg/global/log"
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
	"time"
)

type IMemMapFile interface {
	io.WriteCloser
}

type memMapFileImpl struct {
	ptr []byte
	fd uintptr
	size int64
	fCloser io.Closer
	mutex *sync.Mutex
}

func (mmfi *memMapFileImpl)tryFileLock() (err error) {
	for cnt := 0 ; cnt < 2; cnt ++ {
		err = syscall.Flock(int(mmfi.fd), syscall.LOCK_EX | syscall.LOCK_NB)
		if err == nil {
			break
		}

		if  err != nil && ( !syscall.ENOLCK.Is(err) || !syscall.EWOULDBLOCK.Is(err) ) {
			log.GetLogger().Debug(err.Error())
			break
		}

		time.Sleep(time.Microsecond * 100)
	}
	return 
}

func (mmfi *memMapFileImpl)blockingFileUnLock() error {
	return syscall.Flock(int(mmfi.fd), syscall.LOCK_UN)
}

func (mmfi *memMapFileImpl)Write(b []byte) (n int, err error) {
	if mmfi.ptr == nil || mmfi.mutex == nil {
		log.GetLogger().Debug(fmt.Sprintf("MMAP Write is ptr(%t) or mutex(%t) is nil", mmfi.ptr == nil, mmfi.mutex == nil))
		err = io.ErrNoProgress
		return
	}
	mmfi.mutex.Lock()
	defer mmfi.mutex.Unlock()

	err = mmfi.tryFileLock()
	if err != nil {
		log.GetLogger().Debug(fmt.Sprintf("MMAP Write try file Lock is Error : %s", err.Error()))
		return
	}
	defer mmfi.blockingFileUnLock()
	n += copy(mmfi.ptr[9:],b[9:])
	n += copy(mmfi.ptr[:8], b[:8])
	return
}

func (mmfi *memMapFileImpl)Close() error {
	m := mmfi.mutex
	m.Lock()
	data := mmfi.ptr
	mmfi.mutex = nil
	mmfi.ptr = nil
	m.Unlock()

	mmfi.fCloser.Close()
	return syscall.Munmap(data)
}

func MemMapFileOpen(filename string, size int64) (IMemMapFile, error) {
	f,openErr := os.OpenFile(filename, os.O_CREATE | os.O_RDWR, 0766)
	if openErr != nil {
		return nil, openErr
	}
	//defer f.Close()

	resizeErr := f.Truncate(size)
	if resizeErr != nil {
		log.GetLogger().Debug(resizeErr.Error())
		return nil, resizeErr
	}
	fd := f.Fd()
	data, mmapErr := syscall.Mmap(int(fd), 0, int(size), syscall.PROT_WRITE, syscall.MAP_SHARED);
	if mmapErr != nil {
		log.GetLogger().Debug(mmapErr.Error())
		return nil, mmapErr
	}

	ret := &memMapFileImpl{
		ptr : data,
		size : size,
		mutex: new(sync.Mutex),
		fd : fd,
		fCloser: f,
	}
	return ret,nil
}