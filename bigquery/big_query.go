package bigquery

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"

	"web_scrapper/scrapper" // Import the scrapper package
)

func WriteToBigQuery(pokemonProducts []scrapper.PokemonProduct, projectID, datasetID, tableID, serviceAccountKeyPath string) error {
	ctx := context.Background()

	// Set up BigQuery client with service account key
	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return fmt.Errorf("failed to create BigQuery client: %v", err)
	}

	// Write data to BigQuery
	err = writeToBigQuery(ctx, client, pokemonProducts, datasetID, tableID)
	if err != nil {
		return fmt.Errorf("failed to write data to BigQuery: %v", err)
	}

	return nil
}

func writeToBigQuery(ctx context.Context, client *bigquery.Client, pokemonProducts []scrapper.PokemonProduct, datasetID, tableID string) error {
	// Get the dataset reference
	ds := client.Dataset(datasetID)

	// Check if the dataset exists
	if _, err := ds.Metadata(ctx); err != nil {
		return fmt.Errorf("dataset %s does not exist: %v", datasetID, err)
	}

	// Get the table reference
	table := ds.Table(tableID)

	// Check if the table exists
	if _, err := table.Metadata(ctx); err != nil {
		// If the table does not exist, create a new one
		schema, err := bigquery.InferSchema(scrapper.PokemonProduct{}) // Use scrapper.PokemonProduct
		if err != nil {
			return fmt.Errorf("failed to infer schema for table: %v", err)
		}

		tableMetadata := &bigquery.TableMetadata{
			Schema: schema,
		}

		if err := table.Create(ctx, tableMetadata); err != nil {
			return fmt.Errorf("failed to create table %s: %v", tableID, err)
		}
	}

	// Convert PokemonProduct data to bigquery.ValueSaver
	var rows []*bigquery.StructSaver
	for _, product := range pokemonProducts {
		saver := bigquery.StructSaver{
			Struct:   product,
			InsertID: product.URL,
		}
		rows = append(rows, &saver)
	}

	// Insert the data into the BigQuery table
	if err := table.Inserter().Put(ctx, rows); err != nil {
		return fmt.Errorf("failed to insert data into table: %v", err)
	}

	return nil
}
