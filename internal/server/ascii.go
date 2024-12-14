package server

import (
	"fmt"

	"go.vxn.dev/bbs-go/internal/config"
)

var WelcomeMessage = `
+-------------------------------------------+
|       __    __                            |
|      / /_  / /_  _____      ____  ____    |
|     / __ \/ __ \/ ___/_____/ __ \/ __ \   |
|    / /_/ / /_/ (__  )_____/ /_/ / /_/ /   |
|   /_.___/_.___/____/      \__, /\____/    |
|                          /____/           |
|                                           |
+-------------------------------------------+
vxn-dev bbs-go service (` + config.Version + `)
telnet ` + config.Host + ` ` + fmt.Sprintf("%d", config.Port) + `

`
