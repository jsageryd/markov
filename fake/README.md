# fake (package)

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jsageryd/markov/fake)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](#)

This package uses Markov chains to generate random names.

## Installation
```
go get -u -v github.com/jsageryd/markov/fake
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/jsageryd/markov/fake"
)

func main() {
	fmt.Println("female:", fake.FemaleFirstName(), fake.LastName())
	fmt.Println("  male:", fake.MaleFirstName(), fake.LastName())
}
```
```
female: Marminella Kiwo
  male: Tem Lohanni
```

## Data
The input data for the name generators has been obtained from
[SCB (Statistika Centralbyrån)](https://www.scb.se/).

### Male and female names
- Tilltalsnamn med minst 10 bärare bland folkbokförda 31 december respektive år 1999 - 2017
- http://www.statistikdatabasen.scb.se/goto/sv/ssd/BE0001TNamn10

### Lastnames
- Efternamn med minst 10 bärare bland folkbokförda 31 december respektive år 1999 - 2017
- http://www.statistikdatabasen.scb.se/goto/sv/ssd/BE0001ENamn10

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
