package rule

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func TestRule(t *testing.T) {
}

func TestInspectRune(t *testing.T) {
	word := "def"
	rule := Rule{
		Word:   word,
		Color:  color.FgRed,
		Buffer: &bytes.Buffer{},
	}
	var writer io.Writer
	writer = new(bytes.Buffer)

	text := "abcdefghi"
	reader := bytes.NewBufferString(text)
	bufReader := bufio.NewReader(reader)
	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			if err.Error() != "EOF" {
				t.Fatal(err.Error())
			}
			break
		}

		rule.InspectRune(r, &writer)
		if w, ok := writer.(*bytes.Buffer); ok {
			result := w.String()
			switch r {
			case 'f':
				if !strings.Contains(result, word) {
					t.Fatalf("matched but the word dose not included")
				}
			default:
				if len(result) != 0 {
					t.Fatalf("unexpected word matching (text:%s, word:%s, rune:%c)", text, word, r)
				}
			}
		}

		// reset writer
		writer = new(bytes.Buffer)
	}
}
