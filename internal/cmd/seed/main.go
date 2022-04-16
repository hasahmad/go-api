package main

import (
	"flag"
	"log"
	"os"

	"github.com/hasahmad/go-api/internal/config"
	"github.com/hasahmad/go-api/internal/db/seed"
	"github.com/hasahmad/go-api/internal/helpers"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	args := flag.Args()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := helpers.OpenDbConnection(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("database connection pool established")

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			println(args[1:])
			seed.Execute(cfg, db, args[1:]...)
			os.Exit(0)
		}
	}
}
