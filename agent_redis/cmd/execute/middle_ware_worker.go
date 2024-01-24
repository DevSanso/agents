package main

import (
	"time"
	"os"
	"sync"
	"sync/atomic"
	"math"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/worker"
	"agent_redis/pkg/protos"
	"agent_redis/pkg/format"
)

type MiddleWareWorker struct {
	buf map[int][]byte
	mutex sync.Mutex
	seq int64
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

	formatSeq := atomic.AddInt64(&t.seq, 1)
	if formatSeq >= math.MaxInt {atomic.StoreInt64(&t.seq, 0)}
	
	output, outputErr := format.MakeFormat(int(formatSeq), redisSnapBytes)

	if outputErr != nil{
		return nil, outputErr
	}

	return &worker.WorkerResponse{
		DType: int(protos.SnapFormat_Redis),
		Data: output,
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
