package main

import (
	"net"
	"sync"
)

type Trans interface {
	Send(typ msgType) error
	Receive() (Msg, error)
}

func Listen(addr string) (server Trans, error error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}
	defer tcpListener.Close()

	conn, err := tcpListener.AcceptTCP()
	if err != nil {
		return
	}

	server = &trans{conn: conn, lock: sync.Mutex{}}
	return
}

func Connect(addr string) (client Trans, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	client = &trans{conn: conn, lock: sync.Mutex{}}
	return
}

type trans struct {
	lock   sync.Mutex
	conn   *net.TCPConn
	serial uint32
}

func (t *trans) Send(typ msgType) error {
	t.serial++
	msg, err := NewMsg(t.serial, typ)
	if err != nil {
		return err
	}

	_, err = t.conn.Write(msg.Encode())
	return err
}

func (t *trans) Receive() (Msg, error) {
	return Decode(t.conn)
}
