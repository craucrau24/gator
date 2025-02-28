package cli

import (
	"fmt"

	"github.com/craucrau24/gator/internal/config"
)

type Commands struct {
	handlers map[string]func(*config.State, Command) error
}

func NewCommands() Commands {
	return Commands{handlers: make(map[string]func(*config.State, Command) error)}
}

func (c *Commands) Register(cmd string, handler func(*config.State, Command) error) {
	c.handlers[cmd] = handler
}

func (c *Commands) Run(state *config.State, cmd Command) error {
	fnc, ok := c.handlers[cmd.Cmd]
	if !ok {
		return fmt.Errorf("%s command not found", cmd.Cmd)
	}
	return fnc(state, cmd)
}

func (c *Commands) Init() {
	c.Register("login", handlerLogin)
	c.Register("register", handlerRegister)
	c.Register("reset", handlerReset)
	c.Register("users", handlerUsers)
	c.Register("agg", handlerAgg)
	c.Register("addfeed", middleWareLoggedIn(handlerAddfeed))
	c.Register("feeds", handlerFeeds)
	c.Register("follow", middleWareLoggedIn(handlerFollow))
	c.Register("following", middleWareLoggedIn(handlerFollowing))
}
