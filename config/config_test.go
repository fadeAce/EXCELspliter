package config

import (
	"testing"
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

type ConfigT struct {
	File        string   `yaml:"path"`
	Length      int      `yaml:"length"`
	Head        bool     `yaml:"head"`
	TargetPath  string   `yaml:"target"`
	SourceSheet string   `yaml:"sheet"`
	Output      OutputT  `yaml:"output"`
	Default     DefaultT `yaml:"default"`
}

type OutputT struct {
	Sheet string `yaml:"sheet"`
}

type DefaultT struct {
	Blank bool `yaml:"blank"`
}

func TestConfig(t *testing.T) {
	d, err := ioutil.ReadFile("./yml_test.yaml")
	fmt.Println(err)
	tar := &ConfigT{}
	err = yaml.Unmarshal([]byte(d), &tar)
	fmt.Println(tar)
}
