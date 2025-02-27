package config

import (
	"encoding/json"
	"io"
	"os"
)

var Cfg *Config

type Config struct {
	Service  Service  `json:"service"` 
	Database Database `json:"database"`
	Logging  Logging  `json:"logging"` 
	Debug    bool     `json:"debug"`   
}

type Database struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	Dsn string `json:"dsn"`
}

type Logging struct {
	Level string `json:"level"`
	File  string `json:"file"` 
}

type Service struct {
	Address string `json:"address"`
}

func LoadConfig(filePath string) ( ) {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Error opening config file: " + err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic("Error reading config file: " + err.Error())
	}
	cfg := Config{}
	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		panic("Error parsing config file: " + err.Error())
	}
	Cfg = &cfg
}

