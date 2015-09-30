package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type Config struct {
	Server struct {
		Addr string
		Port string
	}
}

func LoadConfig() (*Config, error) {
	cfg := Config{}
	err := gcfg.ReadFileInto(&cfg, "notification_config.gcfg")
	if err != nil {
		fmt.Println("Error while loading config \n", err)
		return nil, err
	}
	return &cfg, nil
}
