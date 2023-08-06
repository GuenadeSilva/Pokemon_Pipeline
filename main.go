package main

import (
	"flag"
	"web_scrapper/scrapper"

	_ "cloud.google.com/go/bigquery"
)

func main() {
	// Define command-line flag to enable CSV writing
	writeToCSV := flag.Bool("csv", false, "Enable CSV writing")
	flag.Parse()

	// Scrape products with a limit of 5 pages
	pokemonProducts := scrapper.ScrapeProducts(5)

	// Write to CSV if the flag is enabled
	if *writeToCSV {
		scrapper.WriteToCSV(pokemonProducts)
	}
}
