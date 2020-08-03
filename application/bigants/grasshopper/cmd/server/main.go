package main

import (
	"context"
	"log"
	"time"

	"github.com/allscape/bigants/grasshopper/app/auth"

	"github.com/joho/godotenv"

	"github.com/allscape/bigants/grasshopper/app/anthill"
	"github.com/allscape/bigants/grasshopper/app/db"
	"github.com/allscape/bigants/grasshopper/app/depot"
	"github.com/allscape/bigants/grasshopper/app/finder"
	"github.com/allscape/bigants/grasshopper/app/gql"
	"github.com/allscape/bigants/grasshopper/app/secret"
	"github.com/allscape/bigants/grasshopper/app/web"
	"github.com/allscape/bigants/grasshopper/config"
	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

func main() {
	log.Printf("Starting server T+0ms \n")
	started := time.Now()
	ctx := context.Background()

	var err error
	if err = godotenv.Load(); err != nil {
		log.Println(err)
	}

	var (
		dbOptions       = config.DBDatabaseOptions()
		webOptions      = config.WebOptions()
		anthillEndpoint = config.AnthillEndpoint()
		authOptions     = config.AuthOptions()
	)

	var d *db.Database
	if d, err = db.New(ctx, dbOptions); err != nil {
		log.Fatal(err)
	}

	findera := &finder.Adapter{d}
	depota := &depot.Adapter{d}
	anthilla := anthill.NewClient(anthillEndpoint)
	secreta := &secret.Adapter{}

	saps := &sap.Service{findera, findera, depota, anthilla}
	iams := iam.New(findera, depota, secreta)
	auth := auth.New(authOptions, iams)

	gqls, err := gql.New(iams, saps)
	if err != nil {
		log.Fatal(err)
	}

	migrator := db.NewMigrator(d)
	migrator.Log = log.Printf

	if err = migrator.MigrateToLatestVersion(ctx, "app/db/migrations"); err != nil {
		log.Fatal(err)
	}

	log.Printf("Ready to go T+%dms \n", time.Now().Sub(started)/time.Millisecond)
	app := web.New(gqls, iams, auth, webOptions)
	if err = app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
