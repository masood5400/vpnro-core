package main

import (
	"os"

	"github.com/masood5400/vpnro-core/cmd"
)

func main() {
	cmd.ParseCli(os.Args[1:])
}
