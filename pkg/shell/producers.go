package shell

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
)

func NewsProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(NewsHandler)
}

func HelpProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(HelpHandler)
}

func VersionProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(VersionHandler)
}

func FiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(FiveHandler)
}

func DanceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(DanceHandler)
}
