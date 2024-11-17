package main

import (
    "fmt"
    "os"
    "osint-image-scraper/search"
    "osint-image-scraper/scraper"
    "osint-image-scraper/csv"
)

func main() {
    imageUrl := "https://example.com/image.jpg"
    
    results := search.RunReverseImageSearch(imageUrl)
    
    scrapedData := scraper.ScrapeResults(results)
    
    err := csv.ExportToCSV(scrapedData, "output.csv")
    if err != nil {
        fmt.Println("Error exporting to CSV:", err)
        os.Exit(1)
    }
    
    fmt.Println("Data exported successfully to output.csv")
}
