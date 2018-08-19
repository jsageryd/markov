// Package namegen uses Markov chains to generate random names
package namegen

import "strings"

//go:generate go run _generate/main.go

// FemaleFirstName generates a female firstname.
func FemaleFirstName() string {
	return strings.Join(scbFirstNamesFemale.Generate(), "")
}

// MaleFirstName generates a male firstname.
func MaleFirstName() string {
	return strings.Join(scbFirstNamesMale.Generate(), "")
}

// LastName generates a lastname.
func LastName() string {
	return strings.Join(scbLastNames.Generate(), "")
}
