package config

import "os"

type Config struct {
	Port string
	DbUrl string
	JaegerUrl string
	SentryUrl string
	KafkaBroker string
	SomeAppId string
	SomeAppKey string
}

func getWithDefault(envName string, defaultVal string) string {
	val, exist := os.LookupEnv(envName)
	if !exist {
		val = defaultVal
	}
	return val
}

func GetConfig() *Config {
	config := Config{}
	config.Port = getWithDefault("port", "8080")
	config.DbUrl = getWithDefault("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	config.JaegerUrl = getWithDefault("jaeger_url", "http://jaeger:16686")
	config.SentryUrl = getWithDefault("sentry_url", "http://sentry:9000")
	config.KafkaBroker = getWithDefault("kafka_broker", "kafka:9092")
	config.SomeAppId = getWithDefault("some_app_id", "testid")
	config.SomeAppKey = getWithDefault("some_app_key", "testkey")
	return &config
}
