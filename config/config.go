package config

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

var App Application

type Application struct {
	Port     string `yaml:"port"`
	BasePath string `yaml:"base_path"`
	DbUrl    string `yaml:"db_url"`
}

type cfg struct {
	App Application `yaml:"app"`
}

func LoadConfig(path string) error {
	var conf cfg

	yamlFile, err := os.Open(path)
	if err != nil {
		return err
	}

	content, err := io.ReadAll(yamlFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return err
	}

	App = conf.App

	return conf.validate()
}

func (c cfg) validate() error {

	// Port
	if num, err := strconv.Atoi(c.App.Port); err != nil {
		return fmt.Errorf("invalid format app.port: %s", c.App.Port)
	} else if num < 1 {
		return fmt.Errorf("not allowed value app.port: %s", c.App.Port)
	}

	// BasePath
	if _, err := url.Parse(c.App.BasePath); err != nil {
		return fmt.Errorf("invalid format app.baseUrl: %s", c.App.BasePath)
	}

	// PostgreSQL URL
	if _, err := url.Parse(c.App.DbUrl); err != nil {
		return fmt.Errorf("invalid format app.DbUrl: %s", c.App.DbUrl)
	}

	return nil
}
