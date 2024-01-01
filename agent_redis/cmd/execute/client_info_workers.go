package main

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"

	"agent_redis/pkg/global/db"
	"agent_redis/pkg/worker"
	"agent_redis/pkg/protos"
)

type ClientInfoWorker struct {}

func NewClientInfoWorker() *ClientInfoWorker {
	return &ClientInfoWorker{}
}

func (w *ClientInfoWorker)Work(args ...interface{}) (*worker.WorkerResponse, error) {
	ctx,cancelFunc := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancelFunc()

	list, err := db.GetCoreClient().GetClientList(ctx)
	if err != nil {
		return nil, err
	}

	listBytes, err := proto.Marshal(list)
	if err != nil {
		return nil, err
	}

	return &worker.WorkerResponse{
		DType: int(protos.DataFormat_ClientLists),
		Data: listBytes,
	}, nil
}

