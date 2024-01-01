package worker

import (
	"time"
	"context"

	"agent_redis/pkg/global/log"
)
type cmdWorkerInfo struct {
	w IWorker
	interval time.Duration
	ctx context.Context
}
type WorkerManager struct {
	senderWorker IWorker
	middleWareWorker IWorker

	middleChannel chan *WorkerResponse
	sendChannel chan *WorkerResponse

	commandWorkers []cmdWorkerInfo
}
type WorkerManagerBuilder struct {
	wm *WorkerManager
}

func NewWorkerManagerBuilder() *WorkerManagerBuilder {
	b := &WorkerManagerBuilder{wm : new(WorkerManager)}
	b.wm.commandWorkers = make([]cmdWorkerInfo, 0)

	return b
}

func (b *WorkerManagerBuilder)AddCmdWorker(cmdWorker  IWorker, interval time.Duration) *WorkerManagerBuilder {
	b.wm.commandWorkers = append(b.wm.commandWorkers, cmdWorkerInfo{cmdWorker, interval, nil})
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

func (wm *WorkerManager)getNotRunWorkerIndex(idxs []int, output []int) {
	outputIdx := 0
	for _,i := range idxs{
		ele := wm.commandWorkers[i]

		if ele.ctx == nil {
			output[outputIdx] = i
			outputIdx += 1
			continue
		}

		select {
		case <- ele.ctx.Done():
			output[outputIdx] = i
			outputIdx += 1
		default:
		}
	}
}
func (wm *WorkerManager)getIntervalWorkersIndex(output []int) {
	outputIdx := 0
	for i,ele := range wm.commandWorkers {
		if time.Now().UnixMilli() % ele.interval.Milliseconds()  <= 100 {
			output[outputIdx] = i
			outputIdx += 1
		}
	}
}
func runCmdWorker(cmd *cmdWorkerInfo, recv chan <- *WorkerResponse) {
	var cancel context.CancelFunc
	cmd.ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	by,err := cmd.w.Work()
	if err != nil {
		log.GetLogger().Error(err.Error())
		return
	}
	recv <- by
}

func(wm *WorkerManager)cmdWorkerLoop() {
	for {
		intervalWorkers := make([]int, len(wm.commandWorkers))
		willRunWokers := make([]int, len(wm.commandWorkers))

		wm.getIntervalWorkersIndex(intervalWorkers)
		wm.getNotRunWorkerIndex(intervalWorkers, willRunWokers)

		for _,idx := range willRunWokers {
			work := wm.commandWorkers[idx]
			var recv chan <- *WorkerResponse = wm.middleChannel
			go runCmdWorker(&work, recv)
		}
		
		intervalWorkers = nil
		willRunWokers = nil
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
		time.Sleep(time.Microsecond * 5)
	}
}

func (wm *WorkerManager)sendWorkerLoop() {
	for {
		_,err := wm.senderWorker.Work(<-wm.sendChannel)
		if err != nil {
			log.GetLogger().Error(err.Error())
		}
		time.Sleep(time.Microsecond * 10)
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
