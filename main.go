package main

import (
	"flag"
	"log"
	"os"
	"web_scrapper/bigquery"
	"web_scrapper/scrapper"

	"github.com/joho/godotenv"
)

func main() {
	// Define command-line flags
	writeToCSV := flag.Bool("csv", true, "Enable CSV writing")
	writeToBigQuery := flag.Bool("bigquery", false, "Enable BigQuery writing")
	flag.Parse()

	// Scrape products with a limit of 5 pages
	pokemonProducts := scrapper.ScrapeProducts(5)

	// Check if any flags are provided
	if flag.NArg() == 0 {
		// If no flags are provided, use the default behavior (writing to CSV)
		scrapper.WriteToCSV(pokemonProducts)
	} else {
		// Write to CSV if the -csv flag is enabled
		if *writeToCSV {
			scrapper.WriteToCSV(pokemonProducts)
		}

		// Write to BigQuery if the -bigquery flag is enabled
		if *writeToBigQuery {
			errEnv := godotenv.Load()
			if errEnv != nil {
				log.Fatal("Error loading .env file")
			}

			projectID := os.Getenv("PROJECT_ID")
			datasetID := os.Getenv("DATASET_ID")
			tableID := os.Getenv("TABLE_ID")
			serviceAccountKeyPath := os.Getenv("SERVICE_ACCOUNT_KEY_PATH")

			if projectID == "" || datasetID == "" || tableID == "" || serviceAccountKeyPath == "" {
				log.Fatal("PROJECT_ID, DATASET_ID, and SERVICE_ACCOUNT_KEY_PATH must be set in the .env file")
			}
			if serviceAccountKeyPath == "" {
				log.Fatal("Service account key path is required when using BigQuery")
			}

			// Check if the service account key file exists
			if _, err := os.Stat(serviceAccountKeyPath); os.IsNotExist(err) {
				log.Fatalf("Service account key file '%s' does not exist", serviceAccountKeyPath)
			}

			err := bigquery.WriteToBigQuery(pokemonProducts, projectID, datasetID, tableID, serviceAccountKeyPath)
			if err != nil {
				log.Fatalf("Failed to write data to BigQuery: %v", err)
			}
		}
	}
}
