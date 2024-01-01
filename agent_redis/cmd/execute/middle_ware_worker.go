package main

import (
	"time"
	"os"
	"sync"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/worker"
	"agent_redis/pkg/protos"
)

type MiddleWareWorker struct {
	buf map[int][]byte
	mutex sync.Mutex
}

func NewMiddleWareWorker() *MiddleWareWorker {
	return &MiddleWareWorker{
		buf : make(map[int][]byte),
	}
}

func (t *MiddleWareWorker)makeSendData() (*worker.WorkerResponse, error) {
	if len(t.buf) <= 0 {
		return nil, nil
	}

	retSanp := protos.AgentRedisSnap{}

	for key,value := range t.buf {
		temp := &protos.Data{}
		temp.Format = protos.DataFormat(key)
		temp.RawData = value
		retSanp.Datas = append(retSanp.Datas, temp)
	}

	retSanp.UnixEpoch = uint64(time.Now().Unix())
	redisSnapBytes,err := proto.Marshal(&retSanp)

	if err != nil {
		return nil, err
	}

	last := protos.SnapData{}
	last.Format = protos.SnapFormat_Redis
	last.RawSnap = redisSnapBytes

	snapByte, snapErr := proto.Marshal(&last)
	if snapErr != nil {
		return nil, snapErr
	}
	return &worker.WorkerResponse{
		DType: int(protos.SnapFormat_Redis),
		Data: snapByte,
	}, nil
}
func (t *MiddleWareWorker)Work(args ...interface{}) (*worker.WorkerResponse, error) {
	var ret *worker.WorkerResponse = nil
	var retErr error = nil
	
	if args != nil {
		res, ok := args[0].(*worker.WorkerResponse)
		if !ok {
			return nil, os.ErrInvalid
		}
		t.mutex.Lock()
		t.buf[res.DType] = res.Data
		t.mutex.Unlock()
	}

	if time.Now().UnixMilli() % (int64(time.Second) * 4)  < 100 {
		t.mutex.Lock()
		ret, retErr = t.makeSendData()
		t.mutex.Unlock()
	}
	return ret, retErr
}
