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
	version  = "undefined"
	revision = "undefined"
)

var (
	redWords     = kingpin.Flag("red", "Highlight words / red").Short('r').Strings()
	greenWords   = kingpin.Flag("green", "Highlight words / green ").Short('g').Strings()
	yellowWords  = kingpin.Flag("yellow", "Highlight words / yellow").Short('y').Strings()
	blueWords    = kingpin.Flag("blue", "Highlight words / blue").Short('b').Strings()
	magentaWords = kingpin.Flag("magenta", "Highlight words / magenta").Short('m').Strings()
	cyanWords    = kingpin.Flag("cyan", "Highlight words / cyan").Short('c').Strings()
	whiteWords   = kingpin.Flag("white", "Highlight words / white").Short('w').Strings()
	words        = kingpin.Arg("word", "Highlight words / default color (= red)").Strings()

	input      = kingpin.Flag("input", "Input file. If not specified, stdin will be loaded").Short('i').ExistingFile()
	versionOpt = kingpin.Flag("version", "Version information").Short('v').Action(versionOption).Bool()
)

// TODO モードの拡充 (redraw / no-redraw / line buffered)
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
		// defer file.Close()
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

func versionOption(ctx *kingpin.ParseContext) error {
	fmt.Printf("   cring\n Version: %s\nRevision: %s\n", version, revision)
	os.Exit(0)
	return nil
}
