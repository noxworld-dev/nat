package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/noxworld-dev/nat"
)

var (
	fPort  = flag.Int("port", 18590, "port to expose")
	fProto = flag.String("proto", "udp", "network protocol")
	fDesc  = flag.String("desc", "Nox game port", "description for port forward")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	stop, err := nat.Map(context.Background(), []nat.Port{
		{Port: *fPort, Proto: *fProto, Desc: *fDesc},
	})
	if err != nil {
		return err
	}
	defer stop()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	return nil
}
