package csv

import (
    "encoding/csv"
    "os"
    "osint-image-scraper/search"
    "fmt"
)

func ExportToCSV(data []search.SearchResult, fileName string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    writer.Write([]string{"Platform", "URL", "ImageURL", "Metadata"})
    
    for _, result := range data {
        writer.Write([]string{result.Platform, result.URL, result.ImageURL, result.Metadata})
    }
    
    fmt.Println("CSV file written successfully")
    return nil
}

