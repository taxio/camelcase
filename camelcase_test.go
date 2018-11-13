package camelcase

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		options []SplitOption
		hasErr  bool
		want    []string
	}{
		{
			name:    "no input",
			input:   "",
			options: []SplitOption{},
			want:    []string{},
		},
		{
			name:    "lowercase",
			input:   "lowercase",
			options: []SplitOption{},
			want:    []string{"lowercase"},
		},
		{
			name:    "Class",
			input:   "Class",
			options: []SplitOption{},
			want:    []string{"Class"},
		},
		{
			name:    "MyClass",
			input:   "MyClass",
			options: []SplitOption{},
			want: []string{
				"My",
				"Class",
			},
		},
		{
			name:  "MyClass to lower",
			input: "MyClass",
			options: []SplitOption{
				ToLower(true),
			},
			want: []string{
				"my",
				"class",
			},
		},
		{
			name:    "MyC",
			input:   "MyC",
			options: []SplitOption{},
			want: []string{
				"My",
				"C",
			},
		},
		{
			name:    "HTML",
			input:   "HTML",
			options: []SplitOption{},
			want:    []string{"HTML"},
		},
		{
			name:    "PDFLoader",
			input:   "PDFLoader",
			options: []SplitOption{},
			want: []string{
				"PDF",
				"Loader",
			},
		},
		{
			name:    "AString",
			input:   "AString",
			options: []SplitOption{},
			want: []string{
				"A",
				"String",
			},
		},
		{
			name:    "SimpleXMLParser",
			input:   "SimpleXMLParser",
			options: []SplitOption{},
			want: []string{
				"Simple",
				"XML",
				"Parser",
			},
		},
		{
			name:    "vimRPCPlugin",
			input:   "vimRPCPlugin",
			options: []SplitOption{},
			want: []string{
				"vim",
				"RPC",
				"Plugin",
			},
		},
		{
			name:    "GL11Version",
			input:   "GL11Version",
			options: []SplitOption{},
			want: []string{
				"GL",
				"11",
				"Version",
			},
		},
		{
			name:    "99Bottles",
			input:   "99Bottles",
			options: []SplitOption{},
			want: []string{
				"99",
				"Bottles",
			},
		},
		{
			name:    "May5",
			input:   "May5",
			options: []SplitOption{},
			want: []string{
				"May",
				"5",
			},
		},
		{
			name:    "BFG9000",
			input:   "BFG9000",
			options: []SplitOption{},
			want: []string{
				"BFG",
				"9000",
			},
		},
		{
			name:    "BöseÜberraschung",
			input:   "BöseÜberraschung",
			options: []SplitOption{},
			want: []string{
				"Böse",
				"Überraschung",
			},
		},
		{
			name:    "Two  spaces",
			input:   "Two  spaces",
			options: []SplitOption{},
			want: []string{
				"Two",
				"  ",
				"spaces",
			},
		},
		{
			name:    "BadUTF8\xe2\xe2\xa1",
			input:   "BadUTF8\xe2\xe2\xa1",
			options: []SplitOption{},
			want: []string{
				"BadUTF8\xe2\xe2\xa1",
			},
		},
	}

	for _, tt := range tests {
		got := Split(tt.input, tt.options...)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("%s: want: %#v, got: %#v\n", tt.name, tt.want, got)
		}
	}
}
