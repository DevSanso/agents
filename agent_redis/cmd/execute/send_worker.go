package main

import (
	"net"
	"fmt"
	"os"
	"time"
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

func (t *TcpSendWorker)Work(args ...interface{}) ([]byte, error) {
	data, ok := args[0].([]byte)
	if !ok {
		return nil, os.ErrInvalid
	}
	deadline := time.Now().Add(time.Second * 3)
	t.client.SetWriteDeadline(deadline)
	_, err := t.client.Write(data)

	return nil, err
}