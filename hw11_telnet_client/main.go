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
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer client.Close()

	go func() {
		defer cancelFunc()
		if err := client.Send(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}()

	go func() {
		defer cancelFunc()
		if err := client.Receive(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}()

	<-ctx.Done()
}
