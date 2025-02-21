package cli

import "fmt"

type Command struct {
	Cmd  string
	Args []string
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command needs one argument: user name", cmd.Cmd)
	}

	user := cmd.Args[0]
	err := s.Config.SetUser(user)
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", user)
	return nil
}
