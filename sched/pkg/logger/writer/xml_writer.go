package writer

import (
	"sync"
	"io"
	"fmt"
	"time"
)


type xmlLogWriter struct {
	raw io.Writer
	m sync.Mutex
	level string
}

func(xlw *xmlLogWriter)Write(b []byte) (int,error) {
	message := fmt.Sprintf("<Message>%s</Message>", string(b))
	level := fmt.Sprintf("<Level>%s</Level>", xlw.level)

	xlw.m.Lock()
	defer xlw.m.Unlock()

	logtime := fmt.Sprintf("<Logtime>%s</Logtime>", time.Now().Format(time.RFC3339))
	root := fmt.Sprintf(`<Sched>\n
	%s\n
	%s\n
	%s\n
	</Sched>`, logtime, level, message)
	return xlw.raw.Write([]byte(root))

}

func NewXmlWriter(raw io.Writer, level string)  io.Writer {
	return &xmlLogWriter{raw, sync.Mutex{}, level}
}