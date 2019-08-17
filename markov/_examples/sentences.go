package main

import (
	"fmt"
	"strings"

	"github.com/jsageryd/markov/markov"
)

func main() {
	sentences := []string{
		"I am a cat.",
		"As yet I have no name.",
		"I've no idea where I was born.",
	}

	c := markov.NewStringsChain(1, 0)

	for _, s := range sentences {
		c.Feed(strings.Split(s, " "))
	}

	for n := 0; n < 3; n++ {
		fmt.Println(strings.Join(c.Generate(), " "))
	}
}
