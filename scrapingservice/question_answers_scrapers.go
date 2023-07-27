package scrapingservice

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ScrapeForQuestionsAndLinks(madhhab string, page int, c *colly.Collector) {
	var extractedData [][]string
	var wg sync.WaitGroup

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, a *colly.HTMLElement) {
			title := strings.TrimSpace(e.Text)
			url := strings.TrimSpace(a.Attr("href"))
			extractedData = append(extractedData, []string{title, url, madhhab})
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	wg.Add(1)

	var err error

	go func() {
		defer wg.Done()
		err := c.Visit("https://islamqa.org/category/" + madhhab + "/page/" + strconv.Itoa(page))
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	var file *os.File
	if _, err := os.Stat("islamqa_org_fatwas.csv"); os.IsNotExist(err) {
		// Create a new CSV file if it doesn't exist
		file, err = os.Create("islamqa_org_fatwas.csv")
		if err != nil {
			log.Fatal("Error creating CSV file:", err)
		}
		// Write the CSV header
		writer := csv.NewWriter(file)
		err = writer.Write([]string{"title", "url"})
		if err != nil {
			log.Fatal("Error writing CSV header:", err)
		}
		writer.Flush()
	} else {
		// If the file exists, open it in append mode
		file, err = os.OpenFile("islamqa_org_fatwas.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatal("Error opening CSV file:", err)
		}
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range extractedData {
		err = writer.Write(row)
		if err != nil {
			log.Fatal("Error writing CSV data:", err)
		}
	}

	fmt.Println("Loaded questions and answers for " + madhhab + " madhab of page " + strconv.Itoa(page))
}
