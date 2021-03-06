package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
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

var ErrNotConnected = errors.New("not connected")

func (m *MyTelnetClient) Connect() (err error) {
	m.conn, err = net.DialTimeout("tcp", m.address, m.timeout)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "...Connected to "+m.address)
	return nil
}

func (m *MyTelnetClient) Close() (err error) {
	if m.conn == nil {
		return
	}
	if err = m.conn.Close(); err != nil {
		return
	}
	fmt.Fprintln(os.Stderr, "...EOF")
	return nil
}

func (m *MyTelnetClient) Send() error {
	if m.conn == nil {
		fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		return ErrNotConnected
	}
	_, err := io.Copy(m.conn, m.in)
	return err
}

func (m *MyTelnetClient) Receive() error {
	if m.conn == nil {
		fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		return ErrNotConnected
	}
	_, err := io.Copy(m.out, m.conn)
	return err
}
