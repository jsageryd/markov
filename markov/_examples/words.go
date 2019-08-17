package main

import (
	"fmt"
	"strings"

	"github.com/jsageryd/markov/markov"
)

func main() {
	words := []string{
		"albatross",
		"alligator",
		"antelope",
	}

	c := markov.NewStringsChain(2, 0)

	for _, s := range words {
		c.Feed(strings.Split(s, ""))
	}

	for n := 0; n < 3; n++ {
		fmt.Println(strings.Join(c.Generate(), ""))
	}
}
