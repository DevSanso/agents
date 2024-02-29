package main

import (
	"os"
	"time"
	"os/signal"
	"syscall"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/g_var"
)

func main() {

	if(len(os.Args) < 3) {
		panic("Usage: ./execute configPath agent_id")
	}

	configPath := os.Args[1]
	agent_id := os.Args[2]
	

	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic("ReadConfig File Erorr : " + err.Error())
	}


	g_var.InitGlobalVar(agent_id)
	initLogger(cfg)
	initAgentDbClient(cfg)

	ctls, workerErr := initWorkers(cfg)

	if workerErr != nil {
		panic("Init Workers Error : " + workerErr.Error())
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	mainIsRun := true

	for mainIsRun {
		select {
			case <-sigs:
				mainIsRun = false
				continue
			default:
				time.Sleep(time.Second * 10)
		}
	}

	for _, ctl := range ctls {
		ctl.Close()
	}
}