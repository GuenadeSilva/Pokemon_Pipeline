package scrapper

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToCSV(pokemonProducts []PokemonProduct) {
	// Check if there's no data to write
	if len(pokemonProducts) == 0 {
		log.Println("No data to write to CSV.")
		return
	}

	fileName := "products.csv"

	// Check if the file already exists
	if _, err := os.Stat(fileName); err == nil {
		// File exists, so delete it
		err := os.Remove(fileName)
		if err != nil {
			log.Fatalf("Failed to remove existing CSV file: %v", err)
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}
		writer.Write(record)
	}
}
