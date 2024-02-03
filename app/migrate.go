package app

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func (a *App) Migrate() {
	driver, err := postgres.WithInstance(a.DB, &postgres.Config{})
	if err != nil {
		log.Println("postgres instance err: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/",
		"youtubestore", driver)
	log.Println("db instance: ", m)
	if err != nil {
		log.Println("new with db instance error: ", err)
	}
	if err := m.Steps(2); err != nil {
		log.Println("error in steps: ", err)
	}
}

