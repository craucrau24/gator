package main

import (
	"fmt"

	"github.com/craucrau24/gator/internal/config"
)

func main() {
	cfgData, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("database url: %s\n", cfgData.DbUrl)
}
