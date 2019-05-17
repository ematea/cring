package cring

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/ematea/cring/rule"
)

// Coloring colors words according to rules
func Coloring(rules []rule.Rule, reader *io.Reader, writer *io.Writer) error {
	// force enable color if the output destination is tty
	color.NoColor = false

	bufReader := bufio.NewReader(*reader)
	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			if err.Error() != "EOF" {
				return err
			}
			break
		}

		fmt.Fprintf(*writer, "%c", r)
		for _, rule := range rules {
			rule.InspectRune(r, writer)
		}
	}

	return nil
}
