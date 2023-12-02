package main

import (
	"database/sql"
	"time"
	"log"

	"sched/pkg/execute"
)

func runSched(list []NextSched, exec *execute.ScriptExecution) (succList []int, failList []struct { id int; err error}) {
	if len(list) == 0 {
		return
	}

	failList = make([]struct{id int; err error}, 0)
	succList = make([]int, 0)
	
	for _,sched := range list {
		err := exec.Run(sched.SchedName)

		if err != nil {
			log.Println(err.Error())
			failList = append(failList, struct{id int; err error}{id : sched.Id, err : err})
		}else {
			succList = append(succList, sched.Id)
		}
	}

	return
}


func MainLoop(db *sql.DB, scriptExec *execute.ScriptExecution) error {
	for {
		time.Sleep(time.Millisecond * 50)

		schedList,err := DbSelectNextSched(db)		
		if(err != nil) {
			log.Println(err.Error())
			continue
		}

		succ, fail := runSched(schedList, scriptExec)
		
		for _, succId := range succ {
			DbInsertSchedLog(db, succId, "succ", "")
		}

		for _,failed := range fail {
			DbInsertSchedLog(db, failed.id, "fail", failed.err.Error())
		}
		
		succ = nil
		fail = nil		
	}


	return nil
}