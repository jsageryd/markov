// Package fake uses Markov chains to generate random names
package fake

import (
	"strings"
)

//go:generate go run _generate/main.go

// FemaleFirstName generates a female firstname.
func FemaleFirstName() string {
	for {
		name := strings.Join(scbFirstNamesFemale.Generate(), "")
		if !scbFirstNamesFemaleBlacklist.Has([]byte(name)) {
			return name
		}
	}
}

// MaleFirstName generates a male firstname.
func MaleFirstName() string {
	for {
		name := strings.Join(scbFirstNamesMale.Generate(), "")
		if !scbFirstNamesMaleBlacklist.Has([]byte(name)) {
			return name
		}
	}
}

// LastName generates a lastname.
func LastName() string {
	for {
		name := strings.Join(scbLastNames.Generate(), "")
		if !scbLastNamesBlacklist.Has([]byte(name)) {
			return name
		}
	}
}
