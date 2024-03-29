# fake (package)

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jsageryd/markov/fake)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/jsageryd/markov/blob/master/LICENSE)

This package uses Markov chains to generate random names. Names that exist in
the input data are never returned.

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
