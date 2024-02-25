package worker

import (
	"time"
	"fmt"

	"agent_redis/pkg/global/log"
)


type WorkerManager struct {
	senderWorker IWorker
	middleWareWorker IWorker

	middleChannel chan *WorkerResponse
	sendChannel chan *WorkerResponse

	commandWorkers []*cmdWorkerInfo
}
type WorkerManagerBuilder struct {
	wm *WorkerManager
}

func NewWorkerManagerBuilder() *WorkerManagerBuilder {
	b := &WorkerManagerBuilder{wm : new(WorkerManager)}
	b.wm.commandWorkers = make([]*cmdWorkerInfo, 0)

	return b
}

func (b *WorkerManagerBuilder)AddCmdWorker(cmdWorker  IWorker, interval time.Duration) *WorkerManagerBuilder {
	obj := newCmdWorkerInfo(cmdWorker, interval)
	b.wm.commandWorkers = append(b.wm.commandWorkers, obj)
	return b
}

func (b *WorkerManagerBuilder)SendWorker(w IWorker) *WorkerManagerBuilder {
	b.wm.senderWorker = w
	return b
}

func (b *WorkerManagerBuilder)MiddleWareWorker(w IWorker) *WorkerManagerBuilder {
	b.wm.middleWareWorker = w
	return b
}

func (b *WorkerManagerBuilder)Build() *WorkerManager {
	wm := b.wm
	b.wm = nil
	
	wm.middleChannel = make(chan *WorkerResponse)
	wm.sendChannel = make(chan *WorkerResponse)

	return wm
}

func (wm *WorkerManager)StartAndBlock() {
	wm.mainLoop()
}

func (wm *WorkerManager)getNeedRunWorkersIndex(output []int) {
	outputIdx := 0
	for i,ele := range wm.commandWorkers {
		if  ele.isRunNow() {
			log.GetLogger().Debug(fmt.Sprintf("this worker is run now %s", ele.w.GetName()))
			output[outputIdx] = i
			outputIdx += 1
		}
	}
}

func(wm *WorkerManager)cmdWorkerLoop() {
	for {
		intervalWorkers := make([]int, len(wm.commandWorkers))

		wm.getNeedRunWorkersIndex(intervalWorkers)

		for _,idx := range intervalWorkers {
			work := wm.commandWorkers[idx]
			var recv chan <- *WorkerResponse = wm.middleChannel
			work.runCmdWorker(recv)
		}
		
		intervalWorkers = nil
		time.Sleep(time.Millisecond * 500)
	}
}

func (wm *WorkerManager)middleWorkerLoop() {
	var res *WorkerResponse = nil
	var err error = nil

	for {
		select {
		case snap := <- wm.middleChannel:
			res,err = wm.middleWareWorker.Work(snap)
		default:
			res,err = wm.middleWareWorker.Work(nil)
		}
		if err != nil {
			log.GetLogger().Error(err.Error())
		}

		if res != nil {
			wm.sendChannel <- res
		}

		res = nil
		err = nil
		time.Sleep(time.Millisecond * 1)
	}
}

func (wm *WorkerManager)sendWorkerLoop() {
	for {
		_,err := wm.senderWorker.Work(<-wm.sendChannel)
		if err != nil {
			log.GetLogger().Error(err.Error())
		}
		time.Sleep(time.Microsecond * 500)
	}
}

func (wm *WorkerManager)mainLoop() {
	go wm.cmdWorkerLoop()
	go wm.middleWorkerLoop()
	go wm.sendWorkerLoop()

	for {
		time.Sleep(time.Second * 1)
	}
}
