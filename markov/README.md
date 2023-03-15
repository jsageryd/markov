# markov (package)

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jsageryd/markov/markov)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/jsageryd/markov/blob/master/LICENSE)

This package generates markov string sequences.

## Usage
### Words
```go
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
```
```
albator
antelope
alligatross
```

### Sentences
```go
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
```
```
I am a cat.
As yet I am a cat.
I've no name.
```
