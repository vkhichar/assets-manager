package config

import (
	"fmt"
	"os"
	"strconv"
)

type configs struct {
	appPort          int
	dbConfig         DBConfig
	assetServiceURL  string
	assetServicePort int
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

var config configs

func Init() error {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("config: couldn't covert app_port from string to int: %s", err.Error())
		port = 9000
	}

	config.appPort = port
	config.dbConfig = initDBConfig()
	config.assetServiceURL = os.Getenv("ASSET_SERVICE_URL")
	config.assetServicePort, _ = strconv.Atoi(os.Getenv("ASEET_SERVICE_PORT"))

	return nil
}

func initDBConfig() DBConfig {
	cfg := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("config: couldn't read environment variable for db port: %s", err.Error()))
	}
	cfg.Port = port

	return cfg
}

func GetAppPort() string {
	return strconv.Itoa(config.appPort)
}

func GetDBConfig() DBConfig {
	return config.dbConfig
}

func GetAssetServiceURL() string {
	return config.assetServiceURL
}

func GetAssetServicePort() string {
	return strconv.Itoa(config.assetServicePort)
}
