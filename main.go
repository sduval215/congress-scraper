package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Scrape member bioguide for formatted member name (firstName-lastName) and member id
func generateMemberTextFile() {
	// set URL
	var url string = "https://www.congress.gov/help/field-values/member-bioguide-ids"
	// set colly object
	c := colly.NewCollector()

	c.DetectCharset = true

	// read table values
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		// save data rows to tableData variable
		var tableData string = e.ChildText("td")
		// only handle data that exists
		if len(tableData) > 0 {
			// format tableData to member information and member id
			s := strings.SplitAfter(tableData, ")")
			// save member and id to variables
			member, id := s[0], s[1]

			// TODO: FORMAT MEMBER INFORMATION TO FIRSTNAME-LASTNAME FORMAT
			fmt.Printf("Member: %s \n", member)
			// TODO: SAVE ID TO FILE
			fmt.Printf("Member ID: %s \n", id)
		}
	})

	// error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error reading URL: %s, failed with response: %q \n", r.Request.URL, err)
	})

	// on request handler
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %q \n", r.URL.String())
	})

	// visit URL
	c.Visit(url)
}

func main() {
	generateMemberTextFile()
}
