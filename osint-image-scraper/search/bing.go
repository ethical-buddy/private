package bing

import (
    "osint-image-scraper/search"
    "fmt"
)

func Search(imageURL string) search.SearchResult {
    fmt.Println("Searching Bing for image:", imageURL)
    
    return search.SearchResult{
        Platform: "Bing",
        URL:      "https://bing.com/images/",
        ImageURL: imageURL,
        Metadata: "Sample metadata from Bing",
    }
}
