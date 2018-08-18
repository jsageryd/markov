package markov

import (
	"encoding/gob"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
)

const sep = string(0x00)

// StringsChain is a Markov chain of strings.
type StringsChain struct {
	order       int
	starters    [][]string
	transitions map[string][]string
	rng         *rand.Rand
}

// NewStringsChain returns a new StringsChain of the given order.
func NewStringsChain(order int, randSeed int64) *StringsChain {
	return &StringsChain{
		order:       order,
		transitions: make(map[string][]string),
		rng:         rand.New(rand.NewSource(randSeed)),
	}
}

// Feed feeds the chain with the given sequence of strings.
func (c *StringsChain) Feed(seq []string) {
	if len(seq) <= c.order {
		return
	}
	c.starters = append(c.starters, seq[:c.order])
	seq = append(seq, sep)
	for n := c.order; n < len(seq); n++ {
		key := strings.Join(seq[n-c.order:n], sep)
		c.transitions[key] = append(c.transitions[key], seq[n])
	}
}

// Generate generates a sequence of strings based on those fed.
func (c *StringsChain) Generate() []string {
	if len(c.starters) == 0 {
		return []string{}
	}
	s := c.starters[c.rng.Intn(len(c.starters))]
	for s[len(s)-1] != sep {
		key := strings.Join(s[len(s)-c.order:], sep)
		next := c.transitions[key]
		s = append(s, next[c.rng.Intn(len(next))])
	}
	return s[:len(s)-1]
}

// String returns the chain visualised.
func (c *StringsChain) String() string {
	var b strings.Builder
	for n := range c.starters {
		b.WriteString(fmt.Sprintf("%s\n", c.starters[n]))
	}
	keys := make([]string, 0, len(c.transitions))
	for k := range c.transitions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for n := range keys {
		key := strings.Replace(keys[n], sep, " ", -1)
		val := strings.Replace(strings.Join(c.transitions[keys[n]], " "), sep, " ", -1)
		b.WriteString(fmt.Sprintf("\n[%s] -> [%s]", key, val))
	}
	return b.String()
}

// ExportState exports the state of the chain. The state of the PRNG is not
// exported.
func (c *StringsChain) ExportState(w io.Writer) error {
	s := struct {
		Order       int
		Starters    map[string]int
		Transitions map[string]map[string]int
	}{
		Order:       c.order,
		Starters:    make(map[string]int),
		Transitions: make(map[string]map[string]int),
	}

	for n := range c.starters {
		s.Starters[strings.Join(c.starters[n], sep)]++
	}

	for k, v := range c.transitions {
		if _, ok := s.Transitions[k]; !ok {
			s.Transitions[k] = make(map[string]int)
		}
		for n := range v {
			s.Transitions[k][v[n]]++
		}
	}

	return gob.NewEncoder(w).Encode(&s)
}

// ImportState imports the given state. The state of the PRNG is unchanged.
func (c *StringsChain) ImportState(r io.Reader) error {
	var s struct {
		Order       int
		Starters    map[string]int
		Transitions map[string]map[string]int
	}

	if err := gob.NewDecoder(r).Decode(&s); err != nil {
		return err
	}

	c.order = s.Order
	c.starters = make([][]string, 0)
	c.transitions = make(map[string][]string)

	for k, v := range s.Starters {
		x := strings.Split(k, sep)
		for n := 0; n < v; n++ {
			c.starters = append(c.starters, x)
		}
	}

	for k, v := range s.Transitions {
		for kk, vv := range v {
			for n := 0; n < vv; n++ {
				c.transitions[k] = append(c.transitions[k], kk)
			}
		}
	}

	return nil
}
