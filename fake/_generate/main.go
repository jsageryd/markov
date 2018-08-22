package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jsageryd/markov/blacklist"
	"github.com/jsageryd/markov/markov"
)

func main() {
	scbFirstNamesFemaleState, scbFirstNamesFemaleBlacklistState := generateChainAndBlacklistState("_generate/data/scb-firstnames-female-1999-2017.txt")
	scbFirstNamesMaleState, scbFirstNamesMaleBlacklistState := generateChainAndBlacklistState("_generate/data/scb-firstnames-male-1999-2017.txt")
	scbLastNamesState, scbLastNamesBlacklistState := generateChainAndBlacklistState("_generate/data/scb-lastnames-1999-2017.txt")

	src := []byte(`
// generated by go generate; DO NOT EDIT

package fake

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/jsageryd/markov/blacklist"
	"github.com/jsageryd/markov/markov"
)

const (
	scbFirstNamesFemaleState = "` + string(scbFirstNamesFemaleState) + `"
	scbFirstNamesMaleState   = "` + string(scbFirstNamesMaleState) + `"
	scbLastNamesState        = "` + string(scbLastNamesState) + `"

	scbFirstNamesFemaleBlacklistState = "` + string(scbFirstNamesFemaleBlacklistState) + `"
	scbFirstNamesMaleBlacklistState   = "` + string(scbFirstNamesMaleBlacklistState) + `"
	scbLastNamesBlacklistState        = "` + string(scbLastNamesBlacklistState) + `"
)

var (
	scbFirstNamesFemale *markov.StringsChain
	scbFirstNamesMale   *markov.StringsChain
	scbLastNames        *markov.StringsChain

	scbFirstNamesFemaleBlacklist *blacklist.Blacklist
	scbFirstNamesMaleBlacklist   *blacklist.Blacklist
	scbLastNamesBlacklist        *blacklist.Blacklist
)

func init() {
	newChain := func(r io.Reader) *markov.StringsChain {
		newSeed := func() int64 {
			var seed int64
			binary.Read(rand.Reader, binary.BigEndian, &seed)
			return seed
		}
		chain := markov.NewStringsChain(2, newSeed())
		if err := chain.ImportState(base64.NewDecoder(base64.RawStdEncoding, r)); err != nil {
			panic(fmt.Sprintf("error importing chain state: %v", err))
		}
		return chain
	}

	newBlacklist := func(r io.Reader) *blacklist.Blacklist {
		b := blacklist.New()
		if err := b.ImportState(base64.NewDecoder(base64.RawStdEncoding, r)); err != nil {
			panic(fmt.Sprintf("error importing blacklist state: %v", err))
		}
		return b
	}

	scbFirstNamesFemale = newChain(strings.NewReader(scbFirstNamesFemaleState))
	scbFirstNamesMale = newChain(strings.NewReader(scbFirstNamesMaleState))
	scbLastNames = newChain(strings.NewReader(scbLastNamesState))

	scbFirstNamesFemaleBlacklist = newBlacklist(strings.NewReader(scbFirstNamesFemaleBlacklistState))
	scbFirstNamesMaleBlacklist = newBlacklist(strings.NewReader(scbFirstNamesMaleBlacklistState))
	scbLastNamesBlacklist = newBlacklist(strings.NewReader(scbLastNamesBlacklistState))
}
`)

	b, err := format.Source(src)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("gen.go", b, 0644); err != nil {
		log.Fatal(err)
	}
}

func generateChainAndBlacklistState(filename string) ([]byte, []byte) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	chain := markov.NewStringsChain(2, 0)
	blcklist := blacklist.New()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		chain.Feed(strings.Split(string(line), ""))
		blcklist.Add(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var chainBuf bytes.Buffer
	base64Enc := base64.NewEncoder(base64.RawStdEncoding, &chainBuf)
	if err := chain.ExportState(base64Enc); err != nil {
		log.Fatal(err)
	}
	base64Enc.Close()
	chainState := chainBuf.Bytes()

	var blacklistBuf bytes.Buffer
	base64Enc = base64.NewEncoder(base64.RawStdEncoding, &blacklistBuf)
	if err := blcklist.ExportState(base64Enc); err != nil {
		log.Fatal(err)
	}
	base64Enc.Close()
	blacklistState := blacklistBuf.Bytes()

	return chainState, blacklistState
}