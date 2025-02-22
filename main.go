package main

import (
	"fmt"
	"os"

	"github.com/craucrau24/gator/internal/cli"
	"github.com/craucrau24/gator/internal/config"
)

func main() {
	cfgData, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state := cli.State{Config: &cfgData}
	cmds := cli.NewCommands()
	cmds.Init()

	if len(os.Args) < 2 {
		fmt.Println("gator needs at least 1 argument: command name")
		os.Exit(1)
	}

	cmd := cli.Command{Cmd: os.Args[1], Args: os.Args[2:]}
	err = cmds.Run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
