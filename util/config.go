package util

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Listen         string `json:"listen"`
	MattermostHost string `json:"mattermost_url"`
	Token          string `json:"access_token"`
	Secret         string `json:"secret_token"`
}

func LoadConf(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	if err := json.Unmarshal(b, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
