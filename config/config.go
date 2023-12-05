package config

import (
	"os"
	"strconv"
)

type InitManager interface {
	InitConfigMysql()
	GetMySQLConfig() *mySQLConfig
}

type mySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func (m *mySQLConfig) GetMySQLConfig() *mySQLConfig {
	return m
}

func NewDBConfig() InitManager {
	newConfig := new(mySQLConfig)
	newConfig.InitConfigMysql()
	return newConfig
}

func (m *mySQLConfig) InitConfigMysql() {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	m.Username = os.Getenv("DB_USERNAME")
	m.Password = os.Getenv("DB_PASSWORD")
	m.Host = os.Getenv("DB_HOST")
	m.Port = port
	m.Database = os.Getenv("DB_DATABASE")
}
