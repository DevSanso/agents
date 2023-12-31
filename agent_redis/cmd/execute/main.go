package main

import (
	"os"

	"agent_redis/pkg/global/log"
	"agent_redis/pkg/config"
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

	wBuilder := worker.NewWorkerManagerBuilder()
	wBuilder.SendWorker(sendWorker)

	wm := wBuilder.Build()
	wm.StartAndBlock()
}