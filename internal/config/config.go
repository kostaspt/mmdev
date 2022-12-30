package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Auth struct {
		Jira struct {
			Username string `json:"username"`
			ApiToken string `json:"api_token"`
		} `json:"jira"`
	} `json:"auth"`
}

func New() (*Config, error) {
	p, err := configPath()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(p)
	viper.ReadInConfig()

	var c Config
	err = viper.Unmarshal(&c)
	return &c, err
}

func (c *Config) Save() error {
	p, err := configPath()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}

	j, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err = viper.ReadConfig(bytes.NewBuffer(j)); err != nil {
		return err
	}

	return viper.WriteConfigAs(p)
}

func configPath() (string, error) {
	d, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, ".config", "mmdev", "config.json"), nil
}
