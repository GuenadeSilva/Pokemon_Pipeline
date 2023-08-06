package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"web_scrapper/scrapper" // Import the scrapper package

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func UploadToPostgres(pokemonProducts []scrapper.PokemonProduct) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	createTableStmt := `
		CREATE TABLE IF NOT EXISTS pokemon_products (
			id SERIAL PRIMARY KEY,
			url TEXT,
			image TEXT,
			name TEXT,
			price TEXT
		)
	`
	_, err = db.Exec(createTableStmt)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	// Insert data into the table
	for _, product := range pokemonProducts {
		insertStmt := "INSERT INTO pokemon_products (url, image, name, price) VALUES ($1, $2, $3, $4)"
		_, err = db.Exec(insertStmt, product.URL, product.IMAGE, product.NAME, product.PRICE)
		if err != nil {
			log.Printf("Failed to insert data into the table: %v", err)
		}
	}

	log.Println("Data successfully uploaded to PostgreSQL")
	return nil
}
