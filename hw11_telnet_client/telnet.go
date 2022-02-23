package main

import (
	"errors"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &MyTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		conn:    nil,
	}
}

type MyTelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (m *MyTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", m.address, m.timeout)
	if err != nil {
		return err
	}
	m.conn = conn
	return nil
}

func (m *MyTelnetClient) Close() error {
	if m.conn == nil {
		return nil
	}
	return m.conn.Close()
}

func (m *MyTelnetClient) Send() error {
	if m.conn == nil {
		return errors.New("not connected")
	}
	if _, err := io.Copy(m.conn, m.in); err != nil {
		return err
	}
	return nil
}

func (m *MyTelnetClient) Receive() error {
	if m.conn == nil {
		return errors.New("not connected")
	}
	if _, err := io.Copy(m.out, m.conn); err != nil {
		return err
	}
	return nil
}
