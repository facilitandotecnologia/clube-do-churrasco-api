package persistence

import (
	"io/ioutil"
	"github.com/ghodss/yaml"
)

type DBConfig struct {
	Mongodb struct {
		URI      string `yaml:"uri"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongodb"`
}

func NewDBConfig(configFile string) (*DBConfig, error) {
	config := &DBConfig{}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}