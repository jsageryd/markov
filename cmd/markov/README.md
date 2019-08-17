# markov (cli tool)

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](#)

This is a command-line tool for generating Markov string sequences based on
string input.

## Installation
```
go get -u -v github.com/jsageryd/markov/cmd/markov
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
