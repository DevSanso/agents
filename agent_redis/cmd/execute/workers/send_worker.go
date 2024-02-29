package workers

import (
	"fmt"
	"net"
	"os"
	"time"
	"context"

	"agent_redis/pkg/worker"
	"agent_redis/pkg/ipc"
)

type TcpSendWorker struct {
	client *net.TCPConn
}

func NewTcpSendWorker(ip string, port int) (*TcpSendWorker, error) {
	tempConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}
	conn,ok := tempConn.(*net.TCPConn)
	if !ok {
		return nil, &net.ParseError{}
	}
	return &TcpSendWorker{client : conn}, nil
}

func (t *TcpSendWorker)Close() error {
	return t.client.Close()
}
func (w *TcpSendWorker)GetName() string {
	return "TcpSendWorker"
}
func (t *TcpSendWorker)Work(args context.Context) (*worker.WorkerResponse, error) {
	data, ok := args.Value("DATA").([]byte)
	if !ok {
		return nil, os.ErrInvalid
	}
	deadline := time.Now().Add(time.Second * 3)
	t.client.SetWriteDeadline(deadline)
	_, err := t.client.Write(data)

	return nil, err
}

type MmapSendWorker struct {
	client ipc.IMemMapFile
}

func NewMmapSendWorker(filename string, size int) (*MmapSendWorker, error) {
	c, err := ipc.MemMapFileOpen(filename, int64(size))
	if err != nil {
		return nil, err
	}
	return &MmapSendWorker{
		client: c,
	}, nil
}
func (w *MmapSendWorker)GetName() string {
	return "MmapSendWorker"
}
func (t *MmapSendWorker)Close() error {
	return t.client.Close()
}

func (t *MmapSendWorker)Work(args context.Context) (*worker.WorkerResponse, error) {
	data, ok := args.Value("DATA").([]byte)
	if !ok {
		return nil, os.ErrInvalid
	}
	_,err := t.client.Write(data)

	return nil, err
}