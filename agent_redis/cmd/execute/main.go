package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"agent_redis/pkg/config"
	"agent_redis/pkg/global/g_var"
	g_db "agent_redis/pkg/global/db"
	"agent_redis/pkg/global/log"
)

func main() {

	if len(os.Args) < 3 {
		panic("Usage: ./execute configPath agent_id")
	}

	configPath := os.Args[1]
	agent_id := os.Args[2]

	var server *http.Server = nil

	if len(os.Args) == 4 {
		server = makePprofServer(os.Args[3])
	}

	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		panic("ReadConfig File Erorr : " + err.Error())
	}

	initGoRuntime()

	g_var.InitGlobalVar(agent_id)
	err = initLogger(cfg)
	if err != nil {
		panic("InitLogger Erorr : " + err.Error())
	}
	err = initAgentDbClient(cfg)
	if err != nil {
		panic("InitAgentDBClient Erorr : " + err.Error())
	}

	ctls, workerErr := initWorkers(cfg)

	if workerErr != nil {
		panic("Init Workers Error : " + workerErr.Error())
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	mainIsRun := true

	fmt.Println("Agent is running")

	if server != nil {
		go func() {
			pprofErr := server.ListenAndServe()
			if pprofErr != nil {
				log.GetLogger().Error("Pprof Server Error : " + pprofErr.Error())
				mainIsRun = false
			}
		}()
		fmt.Println("Agent Pprof is running")
	}

	for mainIsRun {
		select {
		case sig := <-sigs:
			log.GetLogger().Info(fmt.Sprintf("Agent %s(%s) is setting stop", agent_id, sig.String()))
			mainIsRun = false
			continue
		default:
			time.Sleep(time.Second * 1)
		}
	}

	for _, ctl := range ctls {
		ctl.Close()
	}

	g_db.GetCoreClientCloser().Close()

	close(sigs)
	if server != nil { server.Shutdown(context.Background()) }

	log.GetLogger().Info(fmt.Sprintf("Agent %s is shutdown", agent_id))
}
