package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func Run() {
	c := colly.NewCollector()

	blogs := make([]Blog, 10) // last blog is useless. could potentially increase the capacity of the blog but not necessary right now.

	// the function inside will be run on every html element that matches
	// the #blogindex selector
	c.OnHTML("#blogindex", func(e *colly.HTMLElement) {
		// title := e.ChildText("a")
		// date := e.ChildText(".date")
		// author := e.ChildText(".author")
		titles := e.ChildTexts("a")
		for i := 0; i < len(titles)-1; i++ {
			blogs[i].Title = titles[i]
		}
		dates := e.ChildTexts(".date")
		for i := 0; i < len(dates)-1; i++ {
			blogs[i].Date = dates[i]
		}
		authors := e.ChildTexts(".author")
		for i := 0; i < len(authors)-1; i++ {
			blogs[i].Author = authors[i]
		}
	})

	err := c.Visit("https://go.dev/blog/")
	if err != nil {
		log.Fatal(err)
	}

	for _, blog := range blogs {
		fmt.Printf("Title: %s - Date: %s - Author: %s\n", blog.Title, blog.Date, blog.Author)
	}
}
