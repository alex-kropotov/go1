package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

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

func GetConfigFromDotEnv() *Config {
	loadErr := godotenv.Load()
	if loadErr != nil {
		// ничего не делаю. Или прочитает из файла переменных, или из установленных переменных среды,
		// или получит значения по дефолту
	}
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

func GetConfigFromYaml() *Config {
	content, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &config
}