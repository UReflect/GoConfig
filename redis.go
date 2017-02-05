package config

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type Redis 		struct {
	Host 		string
	Port 		int
	Password	string
}

func (config Config) Redis() Redis {
	var redis Redis

	if err := json.Unmarshal(config.Components["redis"], &redis); err != nil {
		log.Warningf("Config[Redis] : %s: %s", err.Error(), "missing or wrong 'redis' configuration, ignoring")
	}

	if config.Settings["Docker"].(bool) {
		redis.Host = "redis_api"
	} else if redis.Host == "" {
		redis.Host = "127.0.0.1"
		log.Warningf("Config[Redis] : %s%s", "Missing 'host' configuration, assuming default value: ", redis.Host)
	}
	if redis.Port == 0 {
		redis.Port = 6379
		log.Warning("Config[Redis] : %s%d", "Missing 'port' configuration, assuming default value: ", redis.Port)
	}

	return redis
}

func (r Redis) String() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}