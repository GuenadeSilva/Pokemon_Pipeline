# Web Scrapper and BigQuery Loader

This project is a web scrapper that collects product information from a website and allows you to write the scraped data to CSV or BigQuery.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The web scrapper in this project uses the Colly library to scrape product information from a website. The scraped data is stored in a custom `PokemonProduct` struct.

You have the option to either write the scraped data to a CSV file or load it into a BigQuery dataset using a service account key.

## Installation

To use the web scrapper and BigQuery loader, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the official website: https://golang.org/

1. Clone this repository to your local machine:

git clone https://github.com/your-username/web-scrapper-bigquery.git

2. Change into the project directory:

cd web-scrapper-bigquery

3. Install the required dependencies:

go get -u github.com/gocolly/colly
go get -u cloud.google.com/go/bigquery
go get -u github.com/joho/godotenv

## Usage

To use the web scrapper and BigQuery loader, you can run the `main.go` file with appropriate command-line flags.

go run main.go -csv -bigquery

- Use the `-csv` flag to enable CSV writing. This will create a CSV file named `products.csv` in the project directory with the scraped data.

- Use the `-bigquery` flag to enable BigQuery writing. Before using this option, make sure to set up the environment variables in the `.env` file with the necessary configuration details for BigQuery.

If no flag is selected it defaults to CSV output

## Configuration

To configure the BigQuery writing, you need to set up the `.env` file in the project directory with the following variables:
PROJECT_ID=your-project-id
DATASET_ID=your-dataset-id
TABLE_ID = your_table-id
SERVICE_ACCOUNT_KEY_PATH= service_account_file.json

Replace `your-project-id`, `your-dataset-id`, `your-table-id` and `deel_key.json` with the actual values corresponding to your Google Cloud project, BigQuery dataset, and service account key path, respectively.

## Contributing

If you find any issues with the project or want to contribute to it, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
