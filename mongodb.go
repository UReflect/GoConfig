package config

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type MongoDb	struct {
	Host		string
	Port		int
	DB			string
}

func (config Config) MongoDb() MongoDb {
	var mongoDb MongoDb

	if err := json.Unmarshal(config.Components["mongoDb"], &mongoDb); err != nil {
		log.Warningf("Config[mongoDb] : %s: %s", err.Error(), "Missing or wrong 'mongoDb' configuration, ignoring")
	}

	if mongoDb.Host == "" {
		if config.Settings["Docker"].(bool) {
			mongoDb.Host = "mongo_api"
		} else {
			mongoDb.Host = "127.0.0.1"
		}
		log.Warningf("Config[MongoDb] : %s%s", "Missing 'host' configuration, assuming default value: ", mongoDb.Host)
	}
	if mongoDb.Port == 0 {
		if config.Settings["Docker"].(bool) {
			mongoDb.Port = -1
		} else {
			mongoDb.Port = 9001
			log.Warningf("Config[MongoDb] : %s%d", "Missing 'port' configuration, assuming default value: ", mongoDb.Port)
		}
	}
	if mongoDb.DB == "" {
		mongoDb.DB = "boilerplate"
		log.Warningf("Config[MongoDb] : %s%s%s", "Missing 'db' configuration, assuming default value: ", mongoDb.DB)
	}

	return mongoDb
}

func (p MongoDb) String() string {
	if p.Port != -1 {
		return fmt.Sprintf("mongodb://%s:%d", p.Host, p.Port)
	}
	return p.Host
}