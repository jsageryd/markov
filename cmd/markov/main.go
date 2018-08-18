package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jsageryd/markov/markov"
)

func main() {
	var defaultSeed int64
	binary.Read(rand.Reader, binary.BigEndian, &defaultSeed)

	order := flag.Int("order", 1, "The order of the Markov chain")
	seed := flag.Int64("seed", defaultSeed, "Seed for the random number generator")
	split := flag.String("split", " ", "Token to split strings on (use empty string to split on every char)")
	seq := flag.Int("seq", 3, "Number of sequences to generate")
	quiet := flag.Bool("quiet", false, "Avoids printing configuration to stderr on start")
	exp := flag.Bool("export", false, "Instead of generating sequences, export the chain state to stdout")
	imp := flag.Bool("import", false, "Instead of reading input sequences, read the chain state from stdin")

	flag.Parse()

	if !*quiet {
		fmt.Fprintf(os.Stderr, "order:%d seed:%d split:%q seq:%d\n", *order, *seed, *split, *seq)
	}

	c := markov.NewStringsChain(*order, *seed)

	if *imp {
		if err := c.ImportState(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			c.Feed(strings.Split(scanner.Text(), *split))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *exp {
		if err := c.ExportState(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		for n := 0; n < *seq; n++ {
			if s := c.Generate(); len(s) > 0 {
				fmt.Println(strings.Join(s, *split))
			}
		}
	}
}
