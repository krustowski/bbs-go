package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.vxn.dev/bbs-go/internal/server"
)

var signalsToHandle = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

func main() {

	// Create a cancellable context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create new server.
	s := server.NewServer(ctx)

	// Handle a graceful shutdown/configuration reload.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signalsToHandle...)

	go func() {
		for {
			// Fetch the signal to determine the type.
			sig := <-sigChan

			// Handle SIGHUP => configuration reload.
			if sig == syscall.SIGHUP {
				//s.Reload()
				continue
			}

			// Handle other signals => try graceful shutdown.
			signal.Stop(sigChan)
			cancel()
			break
		}
	}()

	// Run the program as daemon.
	if err := s.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	// Program's very exitpoint.
	fmt.Fprintf(os.Stdout, "Program exit.\n")
	os.Exit(0)
}
