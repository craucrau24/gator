package cli

import (
	"context"

	"github.com/craucrau24/gator/internal/database"
)

func middleWareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(s *State, cmd Command) error {
	fnc := func(s *State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}

	return fnc
}
