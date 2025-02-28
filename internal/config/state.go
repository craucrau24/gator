package config

import (
	"github.com/craucrau24/gator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *Config
}
