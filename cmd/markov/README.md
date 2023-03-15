# markov (cli tool)

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/jsageryd/markov/blob/master/LICENSE)

This is a command-line tool for generating Markov string sequences based on
string input.

## Installation
```
go install github.com/jsageryd/markov/cmd/markov@latest
```

## Usage
`markov -h` outputs usage information.

### Generating sentences
```
$ markov -seed=0 <<EOF
> squirrels are members of the family Sciuridae
> squirrels are indigenous to the Americas, Eurasia, and Africa
> the earliest known squirrels date from the Eocene period
> EOF
order:1 seed:0 split:" " seq:3
squirrels are indigenous to the family Sciuridae
squirrels are members of the family Sciuridae
the earliest known squirrels date from the family Sciuridae
```

### Generating words
```
$ markov -split='' -seed=0 <<EOF
> albatross
> alligator
> antelope
> EOF
order:1 seed:0 split:"" seq:3
antope
atropelbateligatr
alballbantope
```

### More tweaking
```
$ markov -order=2 -seq=6 -quiet -split='' -seed=0 <<EOF
> antepenultimate
> bioluminescent
> sesquipedalian
> EOF
ant
bioluminesquipenultimatepedaliant
sesquipent
ante
sescenultimatepedaliante
bioluminesquipenultimatepedalian
```
