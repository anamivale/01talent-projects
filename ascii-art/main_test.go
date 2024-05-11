package main

import (
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	file, _ := os.ReadFile("standard.txt")
	inputFile := strings.Split(string(file), "\n")
	test1, _ := os.ReadFile("test1.txt")
	test2, _ := os.ReadFile("test2.txt")
	type args struct {
		input     string
		inputFile []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{
				input:     "",
				inputFile: inputFile,
			},
			"",
		},
		{
			"words with new line at the end",
			args{
				input:     "Hello\n",
				inputFile: inputFile,
			},
			string(test1),
		},
		{
			"words with two new line in the middle",
			args{
				input:     "Hello\n\nThere",
				inputFile: inputFile,
			},
			string(test2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsciiArt(tt.args.input, tt.args.inputFile); got != tt.want {
				t.Errorf("AsciiArt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnlyNewLines(t *testing.T) {
	tests := []struct {
		name           string
		sepInputString []string
		want           string
	}{
		{
			"Lines with words",
			[]string{"", "welcome", ""},
			"false",
		},
		{
			"only lines",
			[]string{"", "", ""},
			"\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnlyNewLines(tt.sepInputString); got != tt.want {
				t.Errorf("OnlyNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"Empty", "", ""},
		{"Wrong extension", "sample.png", ""},
		{"Right extension", "standard.txt", "standard.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFileName(tt.fileName); got != tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
