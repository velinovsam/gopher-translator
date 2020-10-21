package translation_handler

import (
	"fmt"
	"strings"
	"testing"
)

func TestTranslateWord(t *testing.T) {

	var expected string
	var got string

	got = translateWord("apple")
	expected = "gapple"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with a vowel. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("xray")
	expected = "gexray"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with xr. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("square")
	expected = "aresquogo"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with consonant + qu. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("aquarium")
	expected = "gaquarium"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with vowel + qu. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("hair")
	expected = "airhogo"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with consonant. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("chair")
	expected = "airchogo"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to translate word starting with consonant sound. Expected: %s, got: %s", expected, got))
	}

	got = translateWord("And")
	expected = "Gand"
	if got != expected {
		t.Error(fmt.Sprintf("Failed to keep word capitalization. Expected: %s, got: %s", expected, got))
	}

}

func TestTranslateSentence(t *testing.T) {

	var got string
	var expected string

	got = translateSentence("Parses		all kinds    of \n whitespace?")
	expected = "Arsespogo gall indskogo gof itespacewhogo?"
	if got != expected {
		t.Error("\n" + got + "\n" + expected)
	}

	got = translateSentence("Doesn't translate shortened words.")
	if strings.Contains(got, "'") {
		t.Error("Failed to skip shortened words")
	}

}