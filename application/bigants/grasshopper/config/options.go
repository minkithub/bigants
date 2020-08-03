package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/allscape/bigants/grasshopper/app/auth"
	"github.com/allscape/bigants/grasshopper/app/db"
	"github.com/allscape/bigants/grasshopper/app/web"
)

func DBDatabaseOptions() db.DatabaseOptions {
	return db.DatabaseOptions{
		Host:     getEnv("POSTGRES_HOST"),
		Port:     getEnv("POSTGRES_PORT"),
		User:     getEnv("POSTGRES_USER"),
		Password: getEnv("POSTGRES_PASSWORD"),
		Name:     getEnv("POSTGRES_DB"),
	}
}

func WebOptions() web.Options {
	return web.Options{
		Port:               getEnv("PORT"),
		CorsAllowedOrigins: strings.Split(getEnv("CORS_ALLOWED_ORIGINS"), ","),
	}
}

func AnthillEndpoint() string {
	return getEnv("ANTHILL_ENDPOINT")
}

func AuthOptions() auth.Options {
	return auth.Options{
		AppName:     getEnv("AUTH_APP_NAME"),
		JWTDuration: time.Duration(mustInt64(getEnv("AUTH_JWT_DURATION"))),
		JWTSecret:   getEnv("AUTH_JWT_SECRET"),
	}
}

func mustInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("[WARNING] Env key is missing: %s", key)
	}
	return val
}
