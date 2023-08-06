package scrapper

import (
	"testing"
)

func TestScrapeProducts(t *testing.T) {
	// Test ScrapeProducts function with a limit of 2 pages
	limit := 2
	pokemonProducts := ScrapeProducts(limit)

	// Check if the number of products is correct
	if len(pokemonProducts) != 10 {
		t.Errorf("Expected %d products, but got %d", 10, len(pokemonProducts))
	}

	// Add more specific tests here if needed
}

// You can add more tests for other functions in the scrapper module if needed
