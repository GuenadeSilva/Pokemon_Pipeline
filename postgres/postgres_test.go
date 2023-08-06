package postgres

import (
	"database/sql"
	"fmt"
	"testing"
	"web_scrapper/scrapper" // Import the scrapper package

	_ "github.com/lib/pq"
)

const (
	testHost     = "localhost"
	testPort     = 5432
	testUser     = "postgres"
	testPassword = "12345"
	testDBName   = "postgres"
)

func TestUploadToPostgres(t *testing.T) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", testHost, testPort, testUser, testPassword, testDBName)

	// Print the connection string for debugging
	fmt.Println("Connection string:", connStr)

	// Connect to the test database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("failed to connect to the test database: %v", err)
	}
	defer db.Close()

	// Drop the existing table if it exists
	_, err = db.Exec("DROP TABLE IF EXISTS pokemon_products")
	if err != nil {
		t.Fatalf("failed to drop table: %v", err)
	}

	// Create the table
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
		t.Fatalf("failed to create table: %v", err)
	}

	// Check the initial number of rows in the table
	initialRowCount, err := getRowCount(db)
	if err != nil {
		t.Fatalf("failed to get initial row count: %v", err)
	}
	fmt.Println("Initial row count:", initialRowCount)

	// Prepare some test data
	pokemonProducts := []scrapper.PokemonProduct{
		{
			URL:   "https://example.com/product/1",
			IMAGE: "https://example.com/image/1.jpg",
			NAME:  "Product 1",
			PRICE: "$10.99",
		},
		{
			URL:   "https://example.com/product/2",
			IMAGE: "https://example.com/image/2.jpg",
			NAME:  "Product 2",
			PRICE: "$19.99",
		},
	}

	// Upload the test data to the test database
	err = UploadToPostgres(pokemonProducts)
	if err != nil {
		t.Fatalf("failed to upload data to PostgreSQL: %v", err)
	}

	// Verify that the correct number of rows was inserted
	expectedRowCount := initialRowCount + len(pokemonProducts)
	actualRowCount, err := getRowCount(db)
	if err != nil {
		t.Fatalf("failed to get row count: %v", err)
	}
	fmt.Println("Expected row count:", expectedRowCount)
	fmt.Println("Actual row count:", actualRowCount)

	if actualRowCount != expectedRowCount {
		t.Errorf("expected %d rows, got %d", expectedRowCount, actualRowCount)
	}

	// Drop all rows from the table for the next test
	_, err = db.Exec("DELETE FROM pokemon_products")
	if err != nil {
		t.Fatalf("failed to delete rows from table: %v", err)
	}
}

func getRowCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pokemon_products").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
