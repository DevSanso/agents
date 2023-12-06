package builder

import (
	"io"
	"log"
	"sched/pkg/logger/writer"
)

const (
	XmlFormat = iota
)

type Builder struct {
	raw io.Writer
	flags int
	format int
}

func (b *Builder)RawWriter(w io.Writer) *Builder {
	b.raw = w
	return b
}

func (b *Builder)Flags(flag int) *Builder {
	b.flags = flag
	return b
}

func (b *Builder)UseFormat(format int) * Builder {
	b.format = format
	return b
}

func (b *Builder)Builder() (debugLog *log.Logger, errLog *log.Logger) {
	var formatDebugWriter io.Writer
	var formatErrWriter io.Writer

	if b.format == XmlFormat {
		formatDebugWriter = writer.NewXmlWriter(b.raw,"DEBUG")
		formatErrWriter = writer.NewXmlWriter(b.raw,"Err")
	}else {
		panic("Logger Builder format not support")
	}

	debugLog = log.New(formatDebugWriter, "", b.flags)
	errLog = log.New(formatErrWriter, "", b.flags)
	
	return 
}