package markov

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestStringsChain(t *testing.T) {
	for n, tc := range []struct {
		order         int
		feed          [][]string
		wantChain     string
		wantSequences [][]string
	}{
		{
			order: 1,
			feed: [][]string{
				{"a", "b", "c"},
				{"b", "c", "d"},
				{"c", "d", "e"},
			},
			wantChain: `
[a]
[b]
[c]

[a] -> [b]
[b] -> [c c]
[c] -> [  d d]
[d] -> [  e]
[e] -> [ ]
			`,
			wantSequences: [][]string{
				{"a", "b", "c", "d", "e"},
				{"b", "c"},
				{"a", "b", "c", "d"},
			},
		},
		{
			order: 2,
			feed: [][]string{
				{"a", "b", "c"},
				{"b", "c", "d"},
				{"c", "d", "e"},
			},
			wantChain: `
[a b]
[b c]
[c d]

[a b] -> [c]
[b c] -> [  d]
[c d] -> [  e]
[d e] -> [ ]
			`,
			wantSequences: [][]string{
				{"a", "b", "c", "d"},
				{"c", "d"},
				{"b", "c", "d"},
			},
		},
		{
			order: 3,
			feed: [][]string{
				{"a", "b", "c", "d"},
				{"b", "c", "d", "e"},
				{"c", "d", "e", "f"},
			},
			wantChain: `
[a b c]
[b c d]
[c d e]

[a b c] -> [d]
[b c d] -> [  e]
[c d e] -> [  f]
[d e f] -> [ ]
			`,
			wantSequences: [][]string{
				{"a", "b", "c", "d", "e"},
				{"c", "d", "e"},
				{"b", "c", "d", "e"},
			},
		},
		{
			order:         1,
			feed:          [][]string{{}},
			wantChain:     ``,
			wantSequences: [][]string{{}},
		},
		{
			order:         1,
			feed:          [][]string{{"a"}, {"b"}, {"c"}},
			wantChain:     ``,
			wantSequences: [][]string{{}},
		},
		{
			order:         2,
			feed:          [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}},
			wantChain:     ``,
			wantSequences: [][]string{{}},
		},
		{
			order:         3,
			feed:          [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"c", "d", "e"}},
			wantChain:     ``,
			wantSequences: [][]string{{}},
		},
	} {
		c := NewStringsChain(tc.order, 0)

		for n := range tc.feed {
			c.Feed(tc.feed[n])
		}

		if got, want := c.String(), strings.TrimSpace(tc.wantChain); got != want {
			t.Errorf("[%d] got:\n%v\n\nwant:\n%v", n, got, want)
		}

		for i := range tc.wantSequences {
			gotSequence := c.Generate()

			if len(gotSequence) != len(tc.wantSequences[i]) {
				t.Errorf("[%d] got sequence %s, want %v", n, gotSequence, tc.wantSequences[i])
				continue
			}

			for j := range tc.wantSequences[i] {
				if gotSequence[j] != tc.wantSequences[i][j] {
					t.Errorf("[%d] got sequence %s, want %v", n, gotSequence, tc.wantSequences[i])
					break
				}
			}
		}
	}
}

func TestStringsChain_ImportExportState(t *testing.T) {
	var buf bytes.Buffer

	exported := NewStringsChain(2, 0)

	exported.Feed([]string{"foo", "bar", "baz"})
	exported.Feed([]string{"bar", "baz", "foo"})
	exported.Feed([]string{"baz", "foo", "bar"})

	if err := exported.ExportState(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	imported := NewStringsChain(0, 0)

	if err := imported.ImportState(&buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := imported.order, exported.order; got != want {
		t.Errorf("imported.order = %d, want %d", got, want)
	}

	if got, want := len(imported.starters), len(exported.starters); got != want {
		t.Errorf("imported %d starters, want %d", got, want)
	}

	if got, want := len(imported.transitions), len(exported.transitions); got != want {
		t.Errorf("imported %d transitions keys, want %d", got, want)
	}
}

func ExampleStringsChain_sentences() {
	sentences := []string{
		"squirrels are members of the family Sciuridae",
		"squirrels are indigenous to the Americas, Eurasia, and Africa",
		"the earliest known squirrels date from the Eocene period",
	}

	c := NewStringsChain(1, 0)

	for _, s := range sentences {
		c.Feed(strings.Split(s, " "))
	}

	for n := 0; n < 3; n++ {
		fmt.Println(strings.Join(c.Generate(), " "))
	}

	// Output:
	// squirrels are indigenous to the family Sciuridae
	// squirrels are members of the family Sciuridae
	// the earliest known squirrels date from the family Sciuridae
}

func ExampleStringsChain_words() {
	words := []string{
		"albatross",
		"alligator",
		"antelope",
	}

	c := NewStringsChain(2, 0)

	for _, s := range words {
		c.Feed(strings.Split(s, ""))
	}

	for n := 0; n < 3; n++ {
		fmt.Println(strings.Join(c.Generate(), ""))
	}

	// Output:
	// albator
	// antelope
	// alligatross
}
