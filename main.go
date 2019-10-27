package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// format member data to firstName-lastName format
func formatMemberData(memberData string) string {
	// initial split of data (remove district)
	s1 := strings.Split(memberData, "(")
	// save rawName to variable
	rawName := s1[0]
	// secondary split of data (isolate first and last name)
	s2 := strings.Split(rawName, ",")
	// set initial format layout
	firstName, lastName := s2[1], s2[0]
	// secondary format for first name if it includes dotted abbreviation
	if strings.Contains(firstName, ".") {
		s3 := strings.Split(firstName, ".")
		firstName = s3[0][:len(s3[0])-2]
	}
	// tertiary format for last name if it includes two last names
	if strings.Contains(lastName, " ") {
		s4 := strings.Split(lastName, " ")
		lastName = s4[0] + "-" + s4[1]
	}
	// return formatted string
	return strings.ToLower(firstName + "-" + lastName)
}

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

			memberName := formatMemberData(member)
			// TODO: FORMAT MEMBER INFORMATION TO FIRSTNAME-LASTNAME FORMAT
			fmt.Printf("Member: %s \n", memberName)
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
