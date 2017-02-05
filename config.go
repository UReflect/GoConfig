package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Components 	map[string]json.RawMessage
	Settings   	map[string]interface{}
	Docker		bool
}

func (config Config) Int(setting string) int {
	src, ok := config.Settings[setting].(float64)
	if !ok {
		return -1
	}
	return int(src)
}

func (config Config) Bool(setting string) bool {
	value, ok := config.Settings[setting].(bool)
	if !ok {
		return false
	}
	return value
}

func Parse(file string) (Config, error) {
	var config Config

	if os.Getenv("ENV") == "DOCKER" {
		config.Docker = true
		log.Info("Docker : %s", "Start in Docker Mode")
	}

	f, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(buf, &config); err != nil {
		return config, err
	}

	return config, nil
}