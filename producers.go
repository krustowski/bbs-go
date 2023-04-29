package main

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
)

func newsProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(newsHandler)
}

func helpProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(helpHandler)
}

func versionProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(versionHandler)
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(fiveHandler)
}

func danceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(danceHandler)
}

