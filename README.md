# Markov

[![Build Status](https://travis-ci.com/jsageryd/markov.svg?branch=master)](https://travis-ci.com/jsageryd/markov)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsageryd/markov)](https://goreportcard.com/report/github.com/jsageryd/markov)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](#)

Markov chains with configurable order and random seed.

## Installation
### Command-line tool
```
go get -u -v github.com/jsageryd/markov/cmd/markov
```

### Package
```
go get -u -v github.com/jsageryd/markov
```

## Usage
### Command-line tool
See the readme at [cmd/markov](cmd/markov).

### Package
```go
package main

import (
	"fmt"
	"strings"

	"github.com/jsageryd/markov"
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

## License
Copyright (c) 2018 Johan Sageryd <j@1616.se>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
