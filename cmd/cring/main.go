package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ematea/cring"
	"github.com/ematea/cring/rule"
)

var (
	version  = ""
	revision = ""
)

var (
	redWords     = kingpin.Flag("red", "highlight word (red)").Short('r').Strings()
	greenWords   = kingpin.Flag("green", "highlight word (green)").Short('g').Strings()
	yellowWords  = kingpin.Flag("yellow", "highlight word (yellow)").Short('y').Strings()
	blueWords    = kingpin.Flag("blue", "highlight word (blue)").Short('b').Strings()
	magentaWords = kingpin.Flag("magenta", "highlight word (magenta)").Short('m').Strings()
	cyanWords    = kingpin.Flag("cyan", "highlight word (cyan)").Short('c').Strings()
	whiteWords   = kingpin.Flag("white", "highlight word (white)").Short('w').Strings()
	words        = kingpin.Arg("word", "highlight word (default color (= red))").Strings()

	input      = kingpin.Flag("input", "input file").Short('i').ExistingFile()
	versionOpt = kingpin.Flag("version", "version information").Short('v').Action(versionAction).Bool()
)

func main() {
	kingpin.Parse()

	var rules rule.Rules
	rules.Append(color.FgRed, redWords)
	rules.Append(color.FgGreen, greenWords)
	rules.Append(color.FgYellow, yellowWords)
	rules.Append(color.FgBlue, blueWords)
	rules.Append(color.FgMagenta, magentaWords)
	rules.Append(color.FgCyan, cyanWords)
	rules.Append(color.FgWhite, whiteWords)
	rules.Append(color.FgRed, words)

	reader, readerDeferCallback, err := getReader(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer readerDeferCallback()

	writer, writerDeferCallback := getWriter()
	defer writerDeferCallback()

	if err := cring.Coloring(rules, &reader, &writer); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getReader(input *string) (io.Reader, func(), error) {
	var reader io.Reader
	var file os.File
	if *input == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*input)
		if err != nil {
			return nil, func() {}, err
		}
		reader = bufio.NewReader(file)
	}
	return reader, func() { file.Close() }, nil
}

func getWriter() (io.Writer, func()) {
	var writer io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		writer = os.Stdout
	} else {
		writer = bufio.NewWriterSize(os.Stdout, 65536)
	}
	return writer, func() {
		if w, ok := writer.(*bufio.Writer); ok {
			w.Flush()
		}
	}
}

func versionAction(ctx *kingpin.ParseContext) error {
	fmt.Printf("cring ")
	if version == "" || revision == "" {
		fmt.Printf("(undefined version)\n")
	} else {
		fmt.Printf("%s / %s\n", version, revision)
	}
	os.Exit(0)
	return nil
}
