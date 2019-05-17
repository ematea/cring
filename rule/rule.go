package rule

import (
	"bytes"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

// Rule is highlight rule
type Rule struct {
	Word   string
	Color  color.Attribute
	Buffer *bytes.Buffer
}

// InspectRune receive rune and check if it matches the rule
func (rule *Rule) InspectRune(r rune, writer *io.Writer) {
	rule.Buffer.WriteRune(r)
	bufferStr := rule.Buffer.String()

	if bufferStr == rule.Word {
		coloredFunc := rule.coloredFprintFunc()
		coloredFunc(
			*writer,
			"%s%s",
			strings.Repeat("\x08", runewidth.StringWidth(bufferStr)),
			bufferStr,
		)
		rule.Buffer.Reset()
	} else if strings.HasPrefix(rule.Word, bufferStr) && len(rule.Word) > len(bufferStr) {
		// do nothing (keep buffer)
	} else {
		rule.Buffer.Reset()
	}
}

func (rule *Rule) coloredFprintFunc() func(w io.Writer, format string, a ...interface{}) {
	return color.New(color.Bold, rule.Color).FprintfFunc()
}
