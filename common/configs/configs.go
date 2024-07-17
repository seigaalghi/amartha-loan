// configs/configs.go
package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Configs struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
}

var AppConfig *Configs

func init() {
	AppConfig = &Configs{
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHost:     getEnv("DB_HOST", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
