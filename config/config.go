package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port       int           `yaml:"port"`
	API        APIConfig     `yaml:"api"`
	DataSource DataSourceCfg `yaml:"datasource"`
}

type APIConfig struct {
	Outer    APIEndpoint `yaml:"outer"`
	Internal APIEndpoint `yaml:"internal"`
}

type APIEndpoint struct {
	Prefix string `yaml:"prefix"`
}

type DataSourceCfg struct {
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

var ServerCfg *ServerConfig

func LoadConfig() error {
	ServerCfg = &ServerConfig{}
	yamlFile, err := ioutil.ReadFile("config/server.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file due to error: %v", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, ServerCfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML file due to error: %v", err)
		return err
	}
	return nil
}
