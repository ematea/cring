package cring

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/ematea/cring/rule"
)

func TestColoring(t *testing.T) {
	var result string
	var err error

	result, err = coloringTestCase("abcdeabcdabcaba", []string{"abc"})
	if err != nil {
		t.Fatal(err)
	}
	if result != "abc[1;31mabc[0mdeabc[1;31mabc[0mdabc[1;31mabc[0maba" {
		t.Fatal("Invalid result / single rule coloring")
	}
	result, err = coloringTestCase("abcdeabcdabcaba", []string{"abc", "ab"})
	if err != nil {
		t.Fatal(err)
	}
	if result != "ab[1;31mab[0mc[1;31mabc[0mdeab[1;31mab[0mc[1;31mabc[0mdab[1;31mab[0mc[1;31mabc[0mab[1;31mab[0ma" {
		t.Fatal("Invalid result / partly duplicated multi rule coloring")
	}
}

func coloringTestCase(text string, words []string) (string, error) {
	var rules rule.Rules
	rules.Append(color.FgRed, &words)

	var reader io.Reader
	reader = bytes.NewBufferString(text)
	var writer io.Writer
	writer = new(bytes.Buffer)

	if err := Coloring(rules, &reader, &writer); err != nil {
		return "", err
	}
	if w, ok := writer.(*bytes.Buffer); ok {
		return w.String(), nil
	}
	return "", errors.New("cannot cast writer to bytes.Buffer")
}

func ExampleColoring() {
	var rules rule.Rules
	rules.Append(color.FgRed, &[]string{"sample"})

	var reader io.Reader
	reader = bytes.NewBufferString("This is sample text.")
	var writer io.Writer
	writer = os.Stdout

	Coloring(rules, &reader, &writer)
	// Output: This is sample[1;31msample[0m text.
}

func ExampleMultibyteTextColoring() {
	var rules rule.Rules
	rules.Append(color.FgRed, &[]string{"サンプル"})

	var reader io.Reader
	reader = bytes.NewBufferString("これはサンプルテキストです。")
	var writer io.Writer
	writer = os.Stdout

	Coloring(rules, &reader, &writer)
	// Output: これはサンプル[1;31mサンプル[0mテキストです。
}
