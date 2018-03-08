package helper

import (
	"log"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type orderApi struct {
	Host		string
	Port		int
}
type config struct {
	OrderApi	orderApi
}

func ReadConfig() config {
	var config config
	absPath, _ := filepath.Abs("config.toml")
	if _, err := toml.DecodeFile(absPath, &config); err != nil {
		log.Fatal(err)
	}
	return config
}