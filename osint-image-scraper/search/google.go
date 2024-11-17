package search  // Change from "package google" to "package search"

import (
    "fmt"
)

func SearchGoogle(imageURL string) SearchResult {
    fmt.Println("Searching Google for image:", imageURL)
    
    return SearchResult{
        Platform: "Google",
        URL:      "https://images.google.com/",
        ImageURL: imageURL,
        Metadata: "Sample metadata from Google",
    }
}

