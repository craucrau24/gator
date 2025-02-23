package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/craucrau24/gator/internal/database"
	"github.com/craucrau24/gator/internal/rss"
	"github.com/google/uuid"
)

type Command struct {
	Cmd  string
	Args []string
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command needs one argument: user name", cmd.Cmd)
	}

	name := cmd.Args[0]
	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", user.Name)
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%s command needs one argument: user name", cmd.Cmd)
	}

	name := cmd.Args[0]
	now := time.Now()
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return err
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", user.Name)
	return nil

}

func handlerUsers(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return nil
	}
	for _, user := range users {
		if s.Config.CurrentUserName == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}

func handlerReset(s *State, cmd Command) error {
	err := s.DB.Reset(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	rss, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return nil
	}
	fmt.Println(rss)
	return nil
}
