package main

import (
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	"github.com/hasahmad/go-api/internal/config"
	_ "github.com/hasahmad/go-api/internal/db/migrations"
	"github.com/hasahmad/go-api/internal/helpers"
	_ "github.com/lib/pq"
)

func main() {
	flags := flag.NewFlagSet("goose", flag.ExitOnError)
	dir := flags.String("dir", ".", "directory with migration files")

	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	cfg, err := config.NewDbConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := goose.OpenDBWithDriver(cfg.Type, helpers.BuildDbConnString(cfg))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
