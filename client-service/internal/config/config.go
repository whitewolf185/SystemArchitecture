package config

import (
	"os"
	"time"
)

const TimeOutImageDownloading = 2 * time.Minute

type configName string

const (
	// DbDsn - ключи для подключения к базе данных
	DbDsn = configName("db_dsn")
	// ListenPort - порт, который слушает сервис
	ListenPort = configName("listen_port")
	// ServerIP - ip, на который swagger делает запрос
	ServerIP = configName("server_ip")
)

func GetValue(cnfg configName) string {
	return os.Getenv(string(cnfg))
}
