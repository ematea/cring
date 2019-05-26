# cring [![Build Status](https://travis-ci.org/ematea/cring.svg?branch=master)](https://travis-ci.org/ematea/cring)

Highlight the contents of the standard input or text file by specifying any color and any word.

## Usage

```bash
# coloring stdin
$ echo "This is sample text." | cring -r sample

# coloring text file
$ echo "This is sample text." >> sample.txt
$ cring -i sample.txt -r sample
```

Refer to `cring --help` for details on how to use it.

## Install

```bash
$ go get github.com/ematea/cring/cmd/cring
```

