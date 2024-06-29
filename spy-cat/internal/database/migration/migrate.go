package database

import (
	"embed"
	"fmt"
	"log"

	"spy-cat/internal/database"
)

// migrations - files with dumping db state
//
//go:embed *.sql
var migrations embed.FS

func RunMigrations(db database.Database) error {

	files, err := migrations.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read sql directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		log.Printf("Setting up '%s' db state", fileName)

		fileContent, err := migrations.ReadFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", fileName, err)
		}

		_, err = db.Exec(string(fileContent))
		if err != nil {
			return fmt.Errorf("failed to execute file %s: %w", fileName, err)
		}
		log.Printf("Successfully executed '%s'", fileName)
	}
	return nil
}
