package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var host, port string
	var timeout time.Duration
	pflag.DurationVarP(&timeout, "timeout", "t", 10*time.Second, "")
	pflag.Parse()

	if host = pflag.Arg(0); host == "" {
		fmt.Fprintln(os.Stderr, errors.New("must specify host"))
		os.Exit(1)
	}

	if port = pflag.Arg(1); port == "" {
		fmt.Fprintln(os.Stderr, errors.New("must specify port"))
		os.Exit(1)
	}

	address := net.JoinHostPort(host, port)
	telnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := telnetClient.Connect(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		if err := telnetClient.Send(); err != nil {
			telnetClient.Close()
			cancelFunc()
			return
		}
	}()

	go func() {
		if err := telnetClient.Receive(); err != nil {
			telnetClient.Close()
			cancelFunc()
			return
		}
	}()

	<-ctx.Done()
}
