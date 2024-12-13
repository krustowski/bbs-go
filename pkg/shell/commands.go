package shell

import (
	"github.com/reiver/go-telnet/telsh"
)

type Command struct {
	Name        string
	Description string
	Producer    telsh.ProducerFunc
}

var Cmds = []Command{
	{
		Name:        "dance",
		Description: "ASCII dance experiment",
		Producer:    DanceProducer,
	},
	{
		Name:        "news",
		Description: "list news",
		Producer:    NewsProducer,
	},
	{
		Name:        "ver",
		Description: "print project version",
		Producer:    VersionProducer,
	},
}
