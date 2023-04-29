package main

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	shellHandler *telsh.ShellHandler = telsh.NewShellHandler()
	version string = "v0.1.2"
)

func versionHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "\n\rbbs-go telnet service")
	oi.LongWriteString(stdout, "\n\rversion: " + version + "\n\r")

	return nil
}

func versionProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(versionHandler)
}


func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	return nil
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(fiveHandler)
}

func danceHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	for i:=0; i<20; i++ {
		oi.LongWriteString(stdout, "\r⠋")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠙")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠹")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠸")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠼")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠴")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠦")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠧")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠇")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠏")
		time.Sleep(50*time.Millisecond)
	}
	oi.LongWriteString(stdout, "\r \r\n")

	return nil
}

func danceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(danceHandler)
}


func main() {
	var addr string = "bbs.savla.int:5555"

	if os.Getenv("PROJECT_VERSION") != "" {
		version = os.Getenv("PROJECT_VERSION")
	}

	log.Printf(" starting bbs-go telnet service (" + version + ")...")

	shellHandler := telsh.NewShellHandler()

	shellHandler.WelcomeMessage = `
    __    __
   / /_  / /_  _____      ____  ____
  / __ \/ __ \/ ___/_____/ __ \/ __ \
 / /_/ / /_/ (__  )_____/ /_/ / /_/ /
/_.___/_.___/____/      \__, /\____/
                       /____/

savla-dev bbs-go telnet service
telnet bbs.savla.int 5555

`

	fmt.Printf(shellHandler.WelcomeMessage)

	// Register the "help" command.
	commandName	:= "help"
	shellHandler.Register(commandName, telsh.Help(shellHandler))

	// Register the "five" command.
	commandName     = "five"
	commandProducer := telsh.ProducerFunc(fiveProducer)
	shellHandler.Register(commandName, commandProducer)

	// Register the "dance" command.
	commandName      = "dance"
	commandProducer  = telsh.ProducerFunc(danceProducer)
	shellHandler.Register(commandName, commandProducer)
	shellHandler.Register("dance", telsh.ProducerFunc(danceProducer))

	server := &telnet.Server{
		Addr: addr,
		Handler: shellHandler,
	}

	// serve the telnet service
	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
