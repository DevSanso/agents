package worker

import (
	"agent_redis/pkg/global/log"
	"context"
	"fmt"
	"time"
)
type WorkerResponse struct{
	DType int
	Data []byte
}

type IWorker interface{
	Work(ctx context.Context) (*WorkerResponse, error)
	GetName() string
}

type WorkerThreadCtl struct {
	threadCtxCancel context.CancelFunc
	worker IWorker
}

type WorkerThreadStartUpArgs struct {
	WorkerTimeOut time.Duration
	WorkerInterval time.Duration
	SendChan chan *WorkerResponse
	RecvChan chan *WorkerResponse

	threadCtx context.Context
	worker IWorker
	tick *time.Ticker
}

func NewWorkerThread(w IWorker,args WorkerThreadStartUpArgs) *WorkerThreadCtl {
	ctl := &WorkerThreadCtl{}

	args.tick = time.NewTicker(args.WorkerInterval)
	args.threadCtx, ctl.threadCtxCancel = context.WithCancel(context.Background())
	args.worker = w
	
	go startNonBlockThread(args)
	return ctl
}

func (w *WorkerThreadCtl) Close() {
	w.threadCtxCancel()
}

func workerArgsValueSet(parent context.Context, key string, ch <- chan *WorkerResponse) context.Context {
	if ch == nil {
		return parent
	}
	
	select {
	case data := <- ch:
		valueCtx := context.WithValue(parent, key, data)
		return valueCtx
	default:
		return nil
	}
}

func startNonBlockThread(args WorkerThreadStartUpArgs) {
	isRun := true

	for isRun {
		select {
		case <-args.threadCtx.Done():
			isRun = false
		default:
		}
		if !isRun {continue}
		
		<-args.tick.C

		ctx := workerArgsValueSet(context.Background(), "DATA", args.RecvChan)
		if ctx == nil {
			continue
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, args.WorkerTimeOut)
		res, workErr := args.worker.Work(timeoutCtx)
		cancel()

		if workErr != nil {
			log.GetLogger().Debug(fmt.Sprintf("Worker[%s] - ERROR - is return Error : %s", args.worker.GetName(), workErr.Error()))
			continue
		}

		if args.SendChan != nil {
			if res != nil {
				args.SendChan <- res
			}else {
				log.GetLogger().Error(fmt.Sprintf("Worker[%s] send data return nil", args.worker.GetName()))
			}
		}

	}
}