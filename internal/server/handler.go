package server

import (
	"fmt"
	"io"
	"net"
	"slices"
	"strings"
	"sync"

	"go.vxn.dev/bbs-go/internal/config"
)

type Handler struct {
	//ctx context.Context
	done   chan struct{}
	conn   net.Conn
	output io.Writer
	wg     *sync.WaitGroup
}

func (h *Handler) Handle() {
	defer h.wg.Done()

	h.logf("< Incoming connection: %s", h.conn.RemoteAddr().String())
	h.conn.Write([]byte(WelcomeMessage))
	h.conn.Write([]byte("> "))

	defer h.conn.Close()

	pktChan := make(chan string)
	errChan := make(chan error)

	h.wg.Add(2)
	go h.route(pktChan, errChan)
	go h.read(pktChan, errChan)

	<-errChan

	h.logf("> Connection closed: %s", h.conn.RemoteAddr().String())
}

func (h *Handler) read(pktChan chan string, errChan chan error) {
	defer h.wg.Done()
	defer h.debugf("Handler: read closed")

	// Prepare packet's byte allocation.
	tmp := make([]byte, 512)

	var (
		err      error
		haltRead bool
	)

	// Run the read loop.
	for {
		// Halt the loop on any read error.
		if haltRead {
			break
		}

		select {
		case <-h.done:
			h.conn.Write([]byte("Shutdown\n"))
			h.logf("> Connection closed (daemon's shutdown): %s", h.conn.RemoteAddr().String())

			if errChan != nil {
				errChan <- nil
				close(errChan)
			}

			return

		case <-errChan:
			return

		default:
			// Read bytes from the remote conterpart.
			if _, err = h.conn.Read(tmp); err != nil {
				haltRead = true
				continue
			}

			// ASCII code for a newline is 10.
			if slices.Contains(tmp, 10) {
				if pktChan != nil {
					pktChan <- string(tmp)
				}
			}
		}
	}

	// Handle error from the read loop.
	if err != nil {
		switch err {
		case io.EOF:
			// End-of-file
			h.debugf("< EOF")

		case err.(net.Error):
			if _, werr := h.conn.Write([]byte("Too slow\n")); werr != nil {
				h.debugf("Write error: %s", werr.Error())
				return
			}

			h.debugf("> Connection closed (read timeout): %s", h.conn.RemoteAddr().String())
			return

		default:
			h.debugf("< Unexpected read error: %s", err.Error())
			return
		}
	}

	if errChan != nil {
		errChan <- err
		close(errChan)
	}
}

func (h *Handler) route(pktChan chan string, errChan chan error) {
	defer h.wg.Done()
	defer h.debugf("Handler: route closed")

	defer func() {
		if errChan != nil {
			errChan <- nil
			close(errChan)
		}
	}()

	for {
		select {
		case <-h.done:
			h.conn.Write([]byte("Shutdown\n"))
			h.logf("> Connection closed (daemon's shutdown): %s", h.conn.RemoteAddr().String())
			return

		case <-errChan:
			return

		default:
			pkt := <-pktChan

			parts := strings.Split(pkt, "\n")

			switch strings.TrimSpace(parts[0]) {
			case "":

			case "exit":
				h.conn.Write([]byte("*** Bye\n\n"))
				return

			default:
				h.debugf("Invalid command")
				h.conn.Write([]byte("*** Invalid command\n\n"))
			}

			h.conn.Write([]byte("> "))
		}
	}
}

// Common debug info logging wrapper function.
func (h *Handler) debugf(format string, args ...interface{}) {
	if h.output == nil || !config.Debug {
		return
	}

	// Prepend the debug info prefix.
	f := fmt.Sprintf("(dbg) %s\n", format)

	// Write to the output according to running configuration.
	fmt.Fprintf(h.output, f, args...)
}

// Common basic logging wrapper function.
func (h *Handler) logf(format string, args ...interface{}) {
	if h.output == nil {
		return
	}

	// Write to the output according to running configuration.
	fmt.Fprintf(h.output, format+"\n", args...)
}
