package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type APIGatewayConfig struct {
	APIPort           int
	EventsGatewayHost string
	EventsGatewayPort int
}

type ConfigType struct {
	API APIGatewayConfig
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	MakeNewConfig()
}

var Config ConfigType

func MakeNewConfig() *ConfigType {
	Config = ConfigType{
		API: APIGatewayConfig{
			APIPort:           getEnvAsInt("API_PORT", 0),
			EventsGatewayPort: getEnvAsInt("GITHUB_API_KEY", 0),
			EventsGatewayHost: getEnv("EVENTS_GATEWAY_HOST", "0.0.0.0"),
		},
	}
	return &Config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
