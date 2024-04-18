package main

import (
	"strings"
	"testing"
)

func TestUppercase(t *testing.T) {
	input := strings.Fields("This is so exciting (up)")
	expected := "This is so EXCITING"

	funcOutput := uppercase(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestLowercase(t *testing.T) {
	input := strings.Fields("This is so Exciting (low)")
	expected := "This is so exciting"

	funcOutput := lowercase(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestCapitalization(t *testing.T) {
	input := strings.Fields("This is so exciting (cap)")
	expected := "This is so Exciting"

	funcOutput := capitalization(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestHexToDec(t *testing.T) {
	input := strings.Fields("This is so exciting, like 1E (hex) wins")
	expected := "This is so exciting, like 30 wins"

	funcOutput := hexToDec(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestBinToDec(t *testing.T) {
	input := strings.Fields("This is so exciting, like 10 (bin) wins")
	expected := "This is so exciting, like 2 wins"

	funcOutput := binToDec(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestCapitalizeNWords(t *testing.T) {
	input := strings.Fields("This is so exciting, like two (cap, 2) wins")
	expected := "This is so exciting, Like Two wins"

	funcOutput := capapitalizeNWords(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestUppercaseNWords(t *testing.T) {
	input := strings.Fields("This is so exciting, like two (up, 2) wins")
	expected := "This is so exciting, LIKE TWO wins"

	funcOutput := uppercaseNWords(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestLowercaseNWords(t *testing.T) {
	input := strings.Fields("This is so exciting, lIke Two (low, 2) wins")
	expected := "This is so exciting, like two wins"

	funcOutput := lowercaseNWords(input)
	output := strings.Join(funcOutput, " ")

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestArticles(t *testing.T) {
	input := "This is so exciting, like two wins before a end whistle"
	expected := "This is so exciting, like two wins before an end whistle"

	output := articles(input)

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}

func TestPunctuating(t *testing.T) {
	input := "This is so exciting, ,lIke Two wins ??"
	expected := "This is so exciting,, lIke Two wins??"

	output := Punctuating(input)

	if output != expected {
		t.Errorf("uppercase(%q) = %q but expected %q", input, output, expected)
	}
}
