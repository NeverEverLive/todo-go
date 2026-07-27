package core_logger

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level 				string 		`envconfig:"LEVEL" required:"true"`
	Folder 				string		`envconfig:"FOLDER" required:"true"`
	FolderPermission	os.FileMode `envconfig:"FOLDER_PERMISSION" default:"0755"`
	FilePermission 		os.FileMode `envconfig:"FILE_PERMISSION" default:"0644"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("LOGGER", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()

	if err != nil {
		err := fmt.Errorf("get logger config: %w", err)
		panic(err)
	}

	return config
}
