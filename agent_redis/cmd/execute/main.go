package main

import (
	"os"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/log"
	"agent_redis/pkg/worker"
)

func main() {
	configPath := os.Args[1]
	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic(err)
	}
	err = log.InitLogger(cfg.LogFilePath)
	if err != nil {
		panic(err)
	}

	sendWorker, workerErr := NewTcpSendWorker(cfg.Sender.Ip, cfg.Sender.Port)
	if workerErr != nil {
		log.GetLogger().Error(workerErr.Error())
		os.Exit(2)
	}
	middleWareWorker := NewMiddleWareWorker()
	clientInfoCmdWorker := NewClientInfoWorker()

	wBuilder := worker.NewWorkerManagerBuilder()
	wBuilder.SendWorker(sendWorker)
	wBuilder.MiddleWareWorker(middleWareWorker)
	wBuilder.AddCmdWorker(clientInfoCmdWorker, time.Second * 3)
	
	wm := wBuilder.Build()
	wm.StartAndBlock()
}