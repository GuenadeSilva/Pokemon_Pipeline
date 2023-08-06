package scrapper

import (
	"os"
	"testing"
)

func TestWriteToCSV(t *testing.T) {
	// Sample data for testing
	pokemonProducts := []PokemonProduct{
		{
			URL:   "https://example.com/pokemon1",
			IMAGE: "https://example.com/images/pokemon1.png",
			NAME:  "Pokemon 1",
			PRICE: "$10",
		},
		{
			URL:   "https://example.com/pokemon2",
			IMAGE: "https://example.com/images/pokemon2.png",
			NAME:  "Pokemon 2",
			PRICE: "$15",
		},
	}

	// Call WriteToCSV to create the CSV file
	WriteToCSV(pokemonProducts)

	// Check if the CSV file is created
	if _, err := os.Stat("products.csv"); os.IsNotExist(err) {
		t.Errorf("Expected CSV file to be created, but it doesn't exist")
	}

	// You can add more specific tests here to check the content of the CSV file if needed

	// Call WriteToCSV again to check overwrite behavior
	WriteToCSV(pokemonProducts)

	// Check if the CSV file is still there (it should be overwritten)
	if _, err := os.Stat("products.csv"); os.IsNotExist(err) {
		t.Errorf("Expected CSV file to be overwritten, but it doesn't exist")
	}

	// Clean up by removing the created CSV file
	err := os.Remove("products.csv")
	if err != nil {
		t.Errorf("Failed to remove CSV file: %v", err)
	}
}

func TestWriteToCSVWithNoData(t *testing.T) {
	// Call WriteToCSV with no data
	WriteToCSV([]PokemonProduct{})

	// Check if the CSV file is created
	if _, err := os.Stat("products.csv"); os.IsExist(err) {
		t.Errorf("Expected CSV file not to be created, but it exists")
	}

	// Clean up by removing the created CSV file (if exists)
	err := os.Remove("products.csv")
	if err != nil {
		t.Errorf("Failed to remove CSV file: %v", err)
	}
}

// You can add more tests for other functions in the csv_writer module if needed
