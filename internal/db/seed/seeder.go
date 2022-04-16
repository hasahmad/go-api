package seed

import (
	"log"
	"reflect"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasahmad/go-api/internal/config"
	"github.com/jmoiron/sqlx"
)

type Seed struct {
	Config config.Config
	DB     *sqlx.DB
	sql    *goqu.Database
}

func Execute(cfg config.Config, db *sqlx.DB, seedMethodNames ...string) {
	s := New(cfg, db)
	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func New(cfg config.Config, db *sqlx.DB) Seed {
	return Seed{
		Config: cfg,
		DB:     db,
		sql:    goqu.New(cfg.DB.Type, db),
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}
