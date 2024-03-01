package workers

import (
	"time"
	"os"
	"sync"
	"sync/atomic"
	"context"
	"math"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/global/log"
	"agent_redis/pkg/worker"
	"agent_redis/pkg/protos"
	"agent_redis/pkg/format"
	"agent_redis/pkg/global/g_var"
)

type MiddleWareWorker struct {
	buf map[int][]byte
	mutex sync.Mutex
	seq uint64
}

func NewMiddleWareWorker() *MiddleWareWorker {
	return &MiddleWareWorker{
		buf : make(map[int][]byte),
	}
}

func (t *MiddleWareWorker)makeSendData() (*worker.WorkerResponse, error) {
	if len(t.buf) <= 0 && !t.isInterval() {
		return &worker.WorkerResponse{
			DType: int(-1),
			Data: nil,
		}, nil
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

	formatSeq := atomic.AddUint64(&t.seq, 1)
	if formatSeq >= math.MaxUint64 {atomic.StoreUint64(&t.seq, 0)}
	
	output, outputErr := format.MakeFormat(formatSeq, g_var.GetID(), redisSnapBytes)

	if outputErr != nil{
		return nil, outputErr
	}

	log.GetLogger().Debug("MiddleWareWorker:Work: make send data")

	return &worker.WorkerResponse{
		DType: int(-1),
		Data: output,
	}, nil
}
func (w *MiddleWareWorker)GetName() string {
	return "MiddleWareWorker"
}

func (w *MiddleWareWorker)isInterval() bool {
	sec := time.Now().Second()
	return sec % 5 == 0
}

func (t *MiddleWareWorker)Work(args context.Context) (*worker.WorkerResponse, error) {
	var ret *worker.WorkerResponse = nil
	var retErr error = nil
	
	value := args.Value("DATA")

	if value != nil {
		res, ok := value.(*worker.WorkerResponse)
		if !ok {
			log.GetLogger().Debug("MiddleWareWorker:Work:args[0] is not *worker.WorkerResponse")
			return nil, os.ErrInvalid
		}
		t.mutex.Lock()
		t.buf[res.DType] = res.Data
		t.mutex.Unlock()
	}

	t.mutex.Lock()
	ret, retErr = t.makeSendData()
	t.mutex.Unlock()

	time.Sleep(time.Millisecond * 100)
	return ret, retErr
}
