# markov (package)

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jsageryd/markov/markov)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](#)

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

## License
Copyright (c) 2018-2019 Johan Sageryd <j@1616.se>

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
