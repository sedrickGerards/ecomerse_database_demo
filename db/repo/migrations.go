package repo

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // File source for migrations
)

// Migrate function applies migrations to the database.
func Migrate(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+absPath,
		dbURL,
	)

	if err != nil {
		return err
	}
	// defer m.Close()

	defer func() {
		sourceErr, databaseErr := m.Close()
		if sourceErr != nil {
			log.Printf("error closing migration source: %v", sourceErr)
		}
		if databaseErr != nil {
			log.Printf("error closing migration database: %v", databaseErr)
		}
	}()

	// defer func() {
	// if err := m.Close(); err != nil {
	//     log.Println("Error closing migration resource:", err)
	// }
	// }()

	// Apply migrations
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// MigrateDown function rolls back migrations from the database.
func MigrateDown(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+absPath,
		dbURL,
	)
	if err != nil {
		return err
	}
	// defer m.Close()
	defer func() {
    sourceErr, dbErr := m.Close()
    if sourceErr != nil {
        log.Println("Error closing source:", sourceErr)
    }
    if dbErr != nil {
        log.Println("Error closing database:", dbErr)
    }
}()


	// Apply migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
