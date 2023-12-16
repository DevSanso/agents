package worker

import (
	"time"
	"context"
)
type cmdWorkerInfo struct {
	w IWorker
	interval time.Duration
	ctx context.Context
}
type WorkerManager struct {
	senderWorker IWorker
	sendChannel chan []byte

	commandWorkers []cmdWorkerInfo
}
type WorkerManagerBuilder struct {
	wm *WorkerManager
}

func NewWorkerManagerBuilder() *WorkerManagerBuilder {
	return &WorkerManagerBuilder{wm : new(WorkerManager)}
}

func (b *WorkerManagerBuilder)AddCmdWorker(cmdWorker  IWorker, interval time.Duration) *WorkerManagerBuilder {
	b.wm.commandWorkers = append(b.wm.commandWorkers, cmdWorkerInfo{cmdWorker, interval, nil})
	return b
}

func (b *WorkerManagerBuilder)SendWorker(w IWorker) *WorkerManagerBuilder {
	b.wm.senderWorker = w
	return b
}

func (b *WorkerManagerBuilder)Build() *WorkerManager {
	wm := b.wm
	b.wm = nil
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
		if float64(time.Now().Second()) / ele.interval.Seconds() <= 0.1 {
			output[outputIdx] = i
			outputIdx += 1
		}
	}
}
func runCmdWorker(cmd *cmdWorkerInfo, recv chan <-[]byte) {
	var cancel context.CancelFunc
	cmd.ctx, cancel = context.WithCancel(context.Background())
	by,err := cmd.w.Work()
	if err != nil {
		panic(err)
	}
	recv <- by
	cancel()
}

func(wm *WorkerManager)cmdWorkerLoop() {
	for {
		intervalWorkers := make([]int, len(wm.commandWorkers))
		willRunWokers := make([]int, len(wm.commandWorkers))

		wm.getIntervalWorkersIndex(intervalWorkers)
		wm.getNotRunWorkerIndex(intervalWorkers, willRunWokers)

		for _,idx := range willRunWokers {
			work := wm.commandWorkers[idx]
			var recv chan <- []byte = wm.sendChannel
			go runCmdWorker(&work, recv)
		}
		
		intervalWorkers = nil
		willRunWokers = nil
		
	}
}

func (wm *WorkerManager)sendWorkerLoop() {
	for {
		wm.senderWorker.Work(<-wm.sendChannel)
		time.Sleep(time.Microsecond * 10)
	}
}

func (wm *WorkerManager)mainLoop() {
	go wm.cmdWorkerLoop()
	go wm.sendWorkerLoop()

	for {
		time.Sleep(time.Microsecond * 10)
	}
}
