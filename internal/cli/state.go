package cli

import (
	"github.com/craucrau24/gator/internal/config"
	"github.com/craucrau24/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *config.Config
}
