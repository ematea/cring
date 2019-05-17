package rule

import (
	"testing"

	"github.com/fatih/color"
)

func TestAppend(t *testing.T) {
	var rules Rules
	rules.Append(color.FgBlue, &[]string{"blue_colored_word"})
	if rules[0].Color != color.FgBlue || rules[0].Word != "blue_colored_word" {
		t.Fatal("valid rule should be appended to rules")
	}

	rules.Append(color.FgCyan, &[]string{"cyan_colored_word", "cyan_colored_word_2nd"})
	if rules[1].Color != color.FgCyan || rules[1].Word != "cyan_colored_word" {
		t.Fatal("valid rules should be appended to rules")
	}
	if rules[2].Color != color.FgCyan || rules[2].Word != "cyan_colored_word_2nd" {
		t.Fatal("valid rules should be appended to rules")
	}
}
