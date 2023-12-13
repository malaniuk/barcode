package env

import (
	"log"
	"os"
)

func GetString(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Cannot parse %v", key)
	}

	return value
}
