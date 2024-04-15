package goreloaded

import (
	"fmt"
	"testing"
)

var TextEditorTest = []struct {
	name     string
	test     string
	expected string
}{
	{"up", "it (up) was the best of times, it was the worst of times (up) ,", "IT was the best of times, it was the worst of TIMES,"},
	{"indexUpp", "it (up, 1) was the best of times, it was the worst of times (up, 5) ,", "IT was the best of times, it WAS THE WORST OF TIMES,"},
	{"low", " it was the spring of hope, IT (low) WAS THE (low) winter of despair.", "it was the spring of hope, it WAS the winter of despair."},
	{"lowIndex", " it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "it was the spring of hope, it was the winter of despair."},
	{"cap", "it was the age of foolishness (cap) ,", "it was the age of Foolishness,"},
	{"indexCap", "it was the age of wisdom, it was the age of foolishness (cap, 6) ,", "it was the age of wisdom, It Was The Age Of Foolishness,"},
	{"hex", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
	{"punctuation", "Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
	{"vowel", "There a aunt was. A amazing rock!", "There an aunt was. An amazing rock!"},
	{"bin", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
}

func TestTextEditor(t *testing.T) {
	for i, v := range TextEditorTest {
		t.Run(v.name, func(t *testing.T) {
			got := TextEditor(v.test)
			expect := v.expected
			if got != expect {
				t.Errorf("\ngot: %s\n expext: %s\n", got, expect)
			} else {
				fmt.Printf("passed test %d: function %s()\n", i+1, v.name)
			}
		})
	}

}
