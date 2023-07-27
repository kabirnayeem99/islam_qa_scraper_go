package main

import (
	"github.com/gocolly/colly"
	"github.com/kabirnayeem99/islam_qa_scraper_go/scrapingservice"
	"math/rand"
	"time"
)

func main() {
	c := colly.NewCollector()

	var madhabs []string
	var pageCount []int

	madhabs = append(madhabs, "hanafi", "maliki", "shafii", "hanbali")
	pageCount = append(pageCount, 100, 19, 90, 13)

	for _, madhab := range madhabs {
		for page := 1; page < 100; page++ {
			ScrapeForQuestionsAndLinks(madhab, page, c)
			randomDelay := time.Duration(rand.Intn(27)) * time.Second
			time.Sleep(randomDelay)
		}
		randomDelay := time.Duration(rand.Intn(187)) * time.Second
		time.Sleep(randomDelay)
	}
}
