package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var conf Config

type Config struct {
	Server   Service        `yaml:"service"`
	Database DatabaseConfig `yaml:"database"`
}

type Service struct {
	ServiceName string `yaml:"name"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Init(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		fmt.Errorf("Error unmarshalling config: %v", err)
	}
}

func GetConfig() *Config {
	return &conf
}
