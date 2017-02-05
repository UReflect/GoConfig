package config

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type Graylog struct {
	Host     string
	Port     int
}

func (config Config) Graylog() Graylog {
	var graylog Graylog

	if err := json.Unmarshal(config.Components["graylog"], &graylog); err != nil {
		log.Warningf("Config[graylog] : %s: %s", err.Error(), "Missing or wrong 'graylog' configuration, ignoring")
	}

	if config.Settings["Docker"].(bool) {
		graylog.Host = "graylog"
	} else if graylog.Host == "" {
		graylog.Host = "127.0.0.1"
		log.Warningf("Config[Graylog] : %s%s%s", "Missing 'host' configuration, assuming default value: ", graylog.Host)
	}
	if graylog.Port == 0 {
		graylog.Port = 12201
		log.Warningf("Config[Graylog] : %s%d", "Missing 'port' configuration, assuming default value: ", graylog.Port)
	}

	return graylog
}

func (p Graylog) String() string {
	return fmt.Sprintf("%s:%d", p.Host, p.Port)
}