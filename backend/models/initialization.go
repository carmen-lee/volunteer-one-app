package models

import (
	"log"

	"github.com/VolunteerOne/volunteer-one-app/backend/database"
)

type Model interface {
}

// Register all models into this table
var tables = []Model{
	&Object{},
<<<<<<< HEAD
	&Organization{},
	&OrgUsers{},
	&Users{},
=======
>>>>>>> 4e9ea3a (Update (#5))
}

func Init() {
	// Create migration for all of our tables
	for _, model := range tables {
		log.Printf("Database Migration -> %T", model)
		if database.GetDatabase().AutoMigrate(model) != nil {
			log.Fatalf("Could not complete database migration.\n")
		}
	}
	log.Printf("Database migration successful.\n")
}
