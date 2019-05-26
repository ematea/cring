# cring [![Build Status](https://travis-ci.org/ematea/cring.svg?branch=master)](https://travis-ci.org/ematea/cring)

## Usage

```bash
# coloring stdin
$ echo "This is sample text." | cring -r sample

# coloring text file
$ echo "This is sample text." >> sample.txt
$ cring -i sample.txt -r sample
```

## Install

```bash
$ go get github.com/ematea/cring/cmd/cring
```

