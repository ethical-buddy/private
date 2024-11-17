package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Struct for holding search results
type SearchResult struct {
	Platform  string
	URL       string
	ImageURL  string
	Metadata  string
}

// Search Google Images for the given image URL
func SearchGoogle(imageURL string) ([]SearchResult, error) {
	// Build the Google search URL
	searchURL := fmt.Sprintf("https://www.google.com/searchbyimage?image_url=%s", imageURL)

	// Make the HTTP request
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []SearchResult

	// Scrape Google Images search result page (adjust the CSS selectors as necessary)
	doc.Find(".rg_meta").Each(func(i int, s *goquery.Selection) {
		result := SearchResult{
			Platform: "Google",
			URL:      s.Find("a").AttrOr("href", ""),
			ImageURL: s.Find("img").AttrOr("src", ""),
			Metadata: s.Text(),
		}
		results = append(results, result)
	})

	return results, nil
}

// Search Bing Images for the given image URL
func SearchBing(imageURL string) ([]SearchResult, error) {
	// Build the Bing search URL
	searchURL := fmt.Sprintf("https://www.bing.com/images/search?q=imgurl:%s&view=detailv2", imageURL)

	// Make the HTTP request
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []SearchResult

	// Scrape Bing Images search result page (adjust the CSS selectors as necessary)
	doc.Find(".item").Each(func(i int, s *goquery.Selection) {
		result := SearchResult{
			Platform: "Bing",
			URL:      s.Find("a").AttrOr("href", ""),
			ImageURL: s.Find("img").AttrOr("src", ""),
			Metadata: s.Text(),
		}
		results = append(results, result)
	})

	return results, nil
}

// Run reverse image search across Google and Bing
func RunReverseImageSearch(imageURL string) []SearchResult {
	var wg sync.WaitGroup
	results := make([]SearchResult, 0)
	resultChan := make(chan []SearchResult)

	// Google reverse image search
	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := SearchGoogle(imageURL)
		if err == nil {
			resultChan <- res
		}
	}()

	// Bing reverse image search
	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err := SearchBing(imageURL)
		if err == nil {
			resultChan <- res
		}
	}()

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from the channels
	for result := range resultChan {
		results = append(results, result...)
	}

	return results
}

// Export search results to CSV
func ExportToCSV(data []SearchResult, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Platform", "URL", "ImageURL", "Metadata"})

	// Write data
	for _, result := range data {
		writer.Write([]string{result.Platform, result.URL, result.ImageURL, result.Metadata})
	}

	return nil
}

func main() {
	// Example image URL for reverse search
	imageURL := "https://example.com/sample-image.jpg"

	// Run reverse image search on the given image URL
	results := RunReverseImageSearch(imageURL)

	// Export results to CSV file
	err := ExportToCSV(results, "results.csv")
	if err != nil {
		fmt.Println("Error exporting to CSV:", err)
		return
	}

	fmt.Println("Results exported successfully to results.csv")
}

