package main

import (
	"flag"
	"log"
	"os"
	"web_scrapper/bigquery"
	"web_scrapper/scrapper"
)

func main() {
	// Define command-line flags
	writeToCSV := flag.Bool("csv", false, "Enable CSV writing")
	writeToBigQuery := flag.Bool("bigquery", false, "Enable BigQuery writing")
	serviceAccountKeyPath := "deel_key.json"
	flag.Parse()

	// Scrape products with a limit of 5 pages
	pokemonProducts := scrapper.ScrapeProducts(5)

	// Write to CSV if the flag is enabled
	if *writeToCSV {
		scrapper.WriteToCSV(pokemonProducts)
	}

	// Write to BigQuery if the flag is enabled
	if *writeToBigQuery {
		if serviceAccountKeyPath == "" {
			log.Fatal("Service account key path is required when using BigQuery")
		}

		// Check if the service account key file exists
		if _, err := os.Stat(serviceAccountKeyPath); os.IsNotExist(err) {
			log.Fatalf("Service account key file '%s' does not exist", serviceAccountKeyPath)
		}

		err := bigquery.WriteToBigQuery(pokemonProducts, "your-project-id", "your-dataset-id", "pokemon_products", serviceAccountKeyPath)
		if err != nil {
			log.Fatalf("Failed to write data to BigQuery: %v", err)
		}
	}
}
