package scraper

import (
    "osint-image-scraper/search"
    "fmt"
)

func ScrapeResults(results []search.SearchResult) []search.SearchResult {
    for _, result := range results {
        fmt.Println("Scraping data from:", result.URL)
        result.Metadata += " | Scraped additional info"
    }
    return results
}
