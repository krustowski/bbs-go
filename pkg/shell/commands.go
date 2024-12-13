package main

import (
	"github.com/reiver/go-telnet/telsh"
)

type command struct {
	name		string
	description	string
	producer	telsh.ProducerFunc
}

var cmds = []command{
	{
		name: "dance",
		description: "ASCII dance experiment",
		producer: danceProducer,
	},
	{
		name: "news",
		description: "list news",
		producer: newsProducer,
	},
	{
		name: "ver",
		description: "print project version",
		producer: versionProducer,
	},
}

