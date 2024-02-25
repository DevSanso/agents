package worker

import (
	"context"
	"sync"
	"time"

	"agent_redis/pkg/global/log"
)

type cmdWorkerInfoContext struct {
	ctx context.Context
	cancelFunc context.CancelFunc

	mutex sync.Mutex
}
func (c *cmdWorkerInfoContext)Lock()  {
	c.mutex.Lock()
}

func (c *cmdWorkerInfoContext)Unlock()  {
	c.mutex.Unlock()
}

func (c *cmdWorkerInfoContext)SettingContext() {
	c.Lock()
	defer c.Unlock()

	c.ctx,c.cancelFunc = context.WithTimeout(context.Background(), time.Millisecond * 1)
}

func (c *cmdWorkerInfoContext)CancelContext() {
	c.Lock()
	defer c.Unlock()
	
	c.cancelFunc()
	c.ctx = nil
	c.cancelFunc = nil
}

type cmdWorkerInfo struct {
	w IWorker
	interval time.Duration
	runningTime time.Duration
	ctx *cmdWorkerInfoContext
}

func newCmdWorkerInfo(w IWorker, interval time.Duration) *cmdWorkerInfo {
	return &cmdWorkerInfo{
		w : w,
		interval : interval,
		runningTime : time.Duration(0),
		ctx : &cmdWorkerInfoContext{nil, nil, sync.Mutex{}},
	}
}

func (cwi *cmdWorkerInfo)isRunNextTime() bool {
	ret := time.Now().Second() - int(cwi.runningTime.Seconds()) >= 1
	return ret
}
func (cwi *cmdWorkerInfo)isRunNilContext() bool {
	ret := cwi.ctx.ctx == nil && cwi.ctx.cancelFunc == nil
	return ret
}
func (cwi *cmdWorkerInfo)isRunInInterval() bool {
	ret := time.Now().UnixMilli() % cwi.interval.Milliseconds()  <= 100
	return ret
}
func (cwi *cmdWorkerInfo)isRunNow() bool {
	cwi.ctx.Lock()
	defer cwi.ctx.Unlock()

	runNil := cwi.isRunNilContext()
	runNextTime := cwi.isRunNextTime()
	runInInterval := cwi.isRunInInterval()

	if runNil && runNextTime && runInInterval {
		return true
	}
	return false
}
func (cwi *cmdWorkerInfo)runCmdWorker(recv chan <- *WorkerResponse) {
	cwi.runningTime = time.Duration(time.Now().UnixMilli())
	cwi.ctx.SettingContext()
	go func() {
		defer cwi.ctx.CancelContext()
		by,err := cwi.w.Work(cwi.ctx.ctx)
		if err != nil {
			log.GetLogger().Error(err.Error())
			return
		}
		recv <- by
	}()
} 