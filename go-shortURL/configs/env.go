package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Ip   string
	Port int

	RedisAddr     string
	RedisPassword string
	RedisDb       int
}

func (c *Config) ParseConfig() {
	file, err := os.Open("./configs/dev.json")
	if err != nil {
		log.Fatalf("Failed to GetEnv, err: %v\n", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(c)
	if err != nil {
		log.Fatalf("Failed to decode json file, err: %v\n", err)
	}
}
