package writer

import (
	"os"
	"io"
)

const (
	DISPLAY_WRITER = iota
	FILE_WRITER
)

func NewRawWriter(target int, args ...string) io.Writer {
	if target == DISPLAY_WRITER {
		return os.Stdout
	}else if target == FILE_WRITER {
		f,err := os.OpenFile(args[0], os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.FileMode(660))
		if err != nil {
			panic(err)
		}
		return f
	}
	panic("newWriter target is not support")
}

