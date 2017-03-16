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

func addDocker(config *Config) {
	if config.Settings == nil {
		config.Settings = make(map[string]interface{})
	}
	if _, ok := config.Settings["Docker"]; !ok {
		config.Settings["Docker"] = false;
		if os.Getenv("ENV") == "DOCKER" {
			config.Settings["Docker"] = true
		}
	}
}

func Parse(file string) (Config, error) {
	var config Config

	addDocker(&config);

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