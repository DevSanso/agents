package workers

import (
	"context"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/global/db"
	"agent_redis/pkg/global/log"
	"agent_redis/pkg/protos"
	"agent_redis/pkg/worker"
)

type CpuInfoWorker struct {}

func NewCpuInfoWorker() *CpuInfoWorker {
	return &CpuInfoWorker{}
}

func (w *CpuInfoWorker)GetName() string {
	return "CpuInfoWorker"
}
func (w *CpuInfoWorker)Work(args context.Context) (*worker.WorkerResponse, error) {
	ctx := args
	list, err := db.GetCoreClient().InfoCpu(ctx)
	if err != nil {
		log.GetLogger().Debug(err.Error())
		return nil, err
	}

	listBytes, err := proto.Marshal(list)
	if err != nil {
		log.GetLogger().Debug(err.Error())
		return nil, err
	}

	return &worker.WorkerResponse{
		DType: int(protos.DataFormat_ClientLists),
		Data: listBytes,
	}, nil
}

