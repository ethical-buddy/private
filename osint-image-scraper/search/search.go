package search

import (
    "sync"
)

type SearchResult struct {
    Platform  string
    URL       string
    ImageURL  string
    Metadata  string
}

func RunReverseImageSearch(imageURL string) []SearchResult {
    var wg sync.WaitGroup
    results := make([]SearchResult, 0)
    resultChan := make(chan SearchResult)
    
    // Google reverse image search
    wg.Add(1)
    go func() {
        defer wg.Done()
        resultChan <- SearchGoogle(imageURL)  // Call SearchGoogle directly
    }()
    
    // Bing reverse image search
    wg.Add(1)
    go func() {
        defer wg.Done()
        resultChan <- SearchBing(imageURL)  // Call SearchBing directly
    }()
    
    // Close channel when all goroutines complete
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    // Collect results
    for result := range resultChan {
        results = append(results, result)
    }
    
    return results
}

