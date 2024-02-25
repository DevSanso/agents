package main

import (
	"context"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/global/db"
	"agent_redis/pkg/global/log"
	"agent_redis/pkg/protos"
	"agent_redis/pkg/worker"
)

type ClientInfoWorker struct {}

func NewClientInfoWorker() *ClientInfoWorker {
	return &ClientInfoWorker{}
}

func (w *ClientInfoWorker)GetName() string {
	return "ClientInfoWorker"
}
func (w *ClientInfoWorker)Work(args ...interface{}) (*worker.WorkerResponse, error) {
	ctx := args[0].(context.Context)
	list, err := db.GetCoreClient().GetClientList(ctx)
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

