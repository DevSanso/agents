package main

import (
	"os"
	"fmt"
	"log"
	"io"
	"database/sql"

	"sched/pkg/conf"
	"sched/pkg/execute"
)

func InitDbConn(c *conf.Configure) (*sql.DB,error) {
	dbConf := c.DBConfig;
	return sql.Open("postgres", 
		fmt.Sprintf("user=%s password='%s' host=%s port=%d dbname=%s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName))

	
}

func InitLogger(c *conf.Configure) error {
	logConf := c.LogConf
	logOutputConf := logConf.Output
	
	var w io.Writer = nil

	switch logOutputConf.OutputType {
	case "file":
		file,fileErr := os.OpenFile(logOutputConf.Var, os.O_APPEND | os.O_CREATE, os.FileMode(600))
		if fileErr != nil {
			return fileErr
		}
		w = file
	default:
		w = os.Stdout
	}

	log.SetOutput(w)

	return nil
}

func main() {
	confPath := os.Args[1]
	config, readConfErr := conf.ReadTomlConfig(confPath)
	
	if readConfErr != nil {
		panic(readConfErr)
	}

	loggerErr := InitLogger(config)
	if loggerErr != nil {
		panic(loggerErr)
	}

	execution,executeErr := execute.NewScriptExecution("")
	if executeErr != nil {
		panic(executeErr)
	}

	db, dbErr := InitDbConn(config)
	if dbErr != nil {
		panic(dbErr)
	}
	
	mainLoopErr := MainLoop(db, execution)
	if mainLoopErr != nil {
		panic(mainLoopErr)
	}
}