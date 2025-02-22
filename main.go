package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/craucrau24/gator/internal/cli"
	"github.com/craucrau24/gator/internal/config"
	"github.com/craucrau24/gator/internal/database"
)

func main() {
	cfgData, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfgData.DbUrl)
	dbQueries := database.New(db)

	state := cli.State{Config: &cfgData, DB: dbQueries}
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
