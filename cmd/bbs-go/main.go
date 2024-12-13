package main

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"log"
	"os"
)

var (
	host         string
	port         string
	shellHandler *telsh.ShellHandler = telsh.NewShellHandler()
	version      string
)

func main() {
	host = os.Getenv("LISTEN_ADDR")
	port = os.Getenv("LISTEN_PORT")
	version = os.Getenv("PROJECT_VERSION")

	log.Printf(" starting bbs-go telnet service (" + version + ")...")

	// override defaults
	shellHandler.Prompt = "$ "
	shellHandler.ExitMessage = "\n\rGoodbye!\n\r"
	shellHandler.WelcomeMessage = `
    __    __
   / /_  / /_  _____      ____  ____
  / __ \/ __ \/ ___/_____/ __ \/ __ \
 / /_/ / /_/ (__  )_____/ /_/ / /_/ /
/_.___/_.___/____/      \__, /\____/
                       /____/

savla-dev bbs-go telnet service (` + version + `)
telnet ` + host + ` ` + port + `

`
	log.Printf(shellHandler.WelcomeMessage)

	// loop over commands from commands.go
	for _, cmd := range cmds {
		commandProducer := telsh.ProducerFunc(cmd.producer)

		shellHandler.Register(cmd.name, telsh.ProducerFunc(commandProducer))
	}

	helpCommandProducer := telsh.ProducerFunc(helpProducer)
	shellHandler.Register("help", telsh.ProducerFunc(helpCommandProducer))

	//shellHandler.Register("help", telsh.Help(shellHandler))

	// construct a server
	server := &telnet.Server{
		//Addr: host + ":" + port,
		Addr:    ":" + port,
		Handler: shellHandler,
		Logger:  logger{},
	}

	// serve the telnet service
	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
