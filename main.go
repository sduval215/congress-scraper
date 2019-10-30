package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// format member data to firstName-lastName format
func returnFormattedName(memberData string) string {
	// split by comma separator
	s1 := strings.Split(memberData, ", ")
	// get firstName and lastName variables from split
	firstName, lastName := s1[1], s1[0]
	// secondary format for first name if it includes dotted abbreviation
	if strings.Contains(firstName, " ") {
		s2 := strings.Split(firstName, " ")
		firstName = s2[0]
	}
	// tertiary format for users with two last names
	if strings.Contains(lastName, " ") {
		s3 := strings.Split(lastName, " ")
		lastName = s3[0] + "-" + s3[1]
	}
	// quaternary format for users with abbreviated first names
	if strings.Contains(firstName, ".") {
		s4 := strings.Split(firstName, ".")
		firstName = s4[0]
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

		// get base member data and member ID
		var rawMemberData string = e.ChildText("td:first-child")
		var memberID string = e.ChildText("td:last-child")

		// only handle data that exists
		if len(rawMemberData) > 0 && len(memberID) > 0 {

			// isolate member Data
			s := strings.SplitAfter(rawMemberData, "(")
			// save member and id to variables
			rawMemberName := s[0][:len(s[0])-2]
			// return formatted member name to variable
			memberName := returnFormattedName(rawMemberName)

			fmt.Printf("Member: %s \n", memberName)
			fmt.Printf("Member ID: %s \n", memberID)

			// TODO: Save data to file
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
