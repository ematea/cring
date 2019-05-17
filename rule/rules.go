package rule

import (
	"bytes"

	"github.com/fatih/color"
)

// Rules is just an array of Rule
type Rules []Rule

// Append append new Rule
func (rules *Rules) Append(color color.Attribute, words *[]string) {
	for _, word := range *words {
		*rules = append(*rules, Rule{
			Word:   word,
			Color:  color,
			Buffer: &bytes.Buffer{},
		})
	}
}
