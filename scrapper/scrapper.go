package scrapper

import (
	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	URL, IMAGE, NAME, PRICE string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ScrapeProducts(limit int) []PokemonProduct {
	var pokemonProducts []PokemonProduct
	var pagesToScrape []string
	pageToScrape := "https://scrapeme.live/shop/page/1/"
	pagesDiscovered := []string{pageToScrape}
	i := 1

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")
		if !contains(pagesToScrape, newPaginationLink) {
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}
		pokemonProduct.URL = e.ChildAttr("a", "href")
		pokemonProduct.IMAGE = e.ChildAttr("img", "src")
		pokemonProduct.NAME = e.ChildText("h2")
		pokemonProduct.PRICE = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
		if i < limit {
			// Scrape the next page
			c.Visit(pagesToScrape[i])
			i++
		}
	})

	c.Visit(pageToScrape)
	return pokemonProducts
}
