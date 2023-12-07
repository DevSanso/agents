package global

import (
	"io"
	"log"
	"sync"

	c "sched/pkg/conf"
	log_build "sched/pkg/logger/builder"
	log_writer "sched/pkg/logger/writer"
	"sched/pkg/execute"

)

var global struct {
	once sync.Once

	config *c.Configure
	debugLog *log.Logger
	errLog *log.Logger
	exection *execute.ScriptExecution

	err error
}

func NewGlobal(path string) error {
	global.once.Do(func() {
		conf, err := c.ReadTomlConfig(path)
		if err != nil { 
			global.err = err
			return  
		}
		execution, execErr := execute.NewScriptExecution(conf.ScriptOption)
		if execErr != nil {
			global.err = execErr
			return
		}

		var raw_w io.Writer
		if conf.LogConf.LogType == c.LogConfDisplayType {
			raw_w = log_writer.NewRawWriter(log_writer.DISPLAY_WRITER)
		}else if conf.LogConf.LogType == c.LogConfFileType{
			raw_w = log_writer.NewRawWriter(log_writer.FILE_WRITER, conf.LogConf.LogVar)
		}else {
			panic("Log Conf Not support Log Type")
		}
		b  := log_build.Builder{}
		
		b.RawWriter(raw_w)
		b.Flags(log.Lshortfile | log.Lmicroseconds)
		b.UseFormat(log_build.XmlFormat)

		debugLog, errLog := b.Build()

		global.exection = execution
		global.config = conf
		global.debugLog = debugLog
		global.errLog = errLog
	
	})
	return global.err
}

func GetConfig() *c.Configure {
	return global.config
}

func DebugLog() *log.Logger {
	return global.debugLog
}

func ErrorLog() *log.Logger {
	return global.errLog
}

func Execution() *execute.ScriptExecution {
	return global.exection
}
