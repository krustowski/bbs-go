package main


import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"io"
	"time"
)



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

func danceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{

	return telsh.PromoteHandlerFunc(danceHandler)
}


func main() {

	shellHandler := telsh.NewShellHandler()

	shellHandler.WelcomeMessage = `
 __          __ ______  _        _____   ____   __  __  ______ 
 \ \        / /|  ____|| |      / ____| / __ \ |  \/  ||  ____|
  \ \  /\  / / | |__   | |     | |     | |  | || \  / || |__   
   \ \/  \/ /  |  __|  | |     | |     | |  | || |\/| ||  __|  
    \  /\  /   | |____ | |____ | |____ | |__| || |  | || |____ 
     \/  \/    |______||______| \_____| \____/ |_|  |_||______|

`


	// Register the "five" command.
	commandName     := "five"
	commandProducer := telsh.ProducerFunc(fiveProducer)

	shellHandler.Register(commandName, commandProducer)



	// Register the "dance" command.
	commandName      = "dance"
	commandProducer  = telsh.ProducerFunc(danceProducer)

	shellHandler.Register(commandName, commandProducer)



	shellHandler.Register("dance", telsh.ProducerFunc(danceProducer))

	addr := ":5555"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}
}
