package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/allscape/bigants/grasshopper/app/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	ctx := context.Background()

	target := flag.String("target", "", "target migration name")
	flag.Parse()

	d, err := db.New(ctx, db.DatabaseOptions{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_DB"),
	})
	if err != nil {
		log.Fatal(err)
	}

	m := db.NewMigrator(d)
	m.Log = log.Printf

	if err = m.MigrateTo(ctx, "./app/db/migrations/", *target); err != nil {
		log.Fatal(err)
	}
}
