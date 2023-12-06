package execute

import (
	"context"
	"os"
	"bytes"
	"os/exec"
	"path"
	"time"
	"fmt"
	"errors"
)

type ScriptExecution struct {
	scriptExec string
	scriptExecOption string
	dir string
}



func NewScriptExecution(cmdOption string) (*ScriptExecution, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &ScriptExecution{scriptExec : "julia", scriptExecOption: cmdOption ,dir : path.Join(pwd, "script") }, nil
}

func (se *ScriptExecution)mappingScriptPath(schedName string) string {
	return path.Join(se.dir, schedName+".ji")
}

func (se *ScriptExecution)Run(name string, processArgs string) error {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second * 55)
	defer cancel()

	cmd := exec.CommandContext(timeoutCtx, 
		fmt.Sprintf("%s %s %s %s",se.scriptExecOption, se.scriptExec, processArgs, se.mappingScriptPath(name)))
	
	buf := make([]byte, 4096)
	cmd.Stderr = bytes.NewBuffer(buf)

	err := cmd.Run().(*exec.ExitError)
	
	if err == nil {
		return nil
	}

	if err.ProcessState.ExitCode() == 2 {
		return errors.New(string(buf))
	}
	
	return err
}