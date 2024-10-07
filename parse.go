package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Compose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image string   `yaml:"image"`
	Ports []string `yaml:"ports"` // its basically {in-port}:{out-port}
}

func openFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func parse(content []byte) Compose {
	print(string(content))
	var compose Compose
	err := yaml.Unmarshal(content, &compose)
	if err != nil {
		panic(err.Error())
	}
	return compose
}
