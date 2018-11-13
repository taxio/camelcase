# Archived project. No maintenance. 

This project is not maintained anymore and is archived. Feel free to fork and
make your own changes if needed. For more detail read my blog post: [Taking an indefinite sabbatical from my projects](https://arslan.io/2018/10/09/taking-an-indefinite-sabbatical-from-my-projects/)

Thanks to everyone for their valuable feedback and contributions.


# CamelCase [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/taxio/camelcase) [![Build Status](https://travis-ci.org/taxio/camelcase.svg?branch=master)](https://travis-ci.org/taxio/camelcase)

CamelCase is a Golang (Go) package to split the words of a camelcase type
string into a slice of words. It can be used to convert a camelcase word (lower
or upper case) into any type of word.

## Splitting rules:

1. If string is not valid UTF-8, return it without splitting as
   single item array.
2. Assign all unicode characters into one of 4 sets: lower case
   letters, upper case letters, numbers, and all other characters.
3. Iterate through characters of string, introducing splits
   between adjacent characters that belong to different sets.
4. Iterate through array of split strings, and if a given string
   is upper case:
   * if subsequent string is lower case:
     * move last character of upper case string to beginning of
       lower case string

## Install

```bash
go get github.com/fatih/camelcase
```

## Usage and examples

```go
splitted := camelcase.Split("GolangPackage")

fmt.Println(splitted[0], splitted[1]) // prints: "Golang", "Package"
```

Both lower camel case and upper camel case are supported. For more info please
check: [http://en.wikipedia.org/wiki/CamelCase](http://en.wikipedia.org/wiki/CamelCase)

Below are some example cases:

```
"" =>                     []
"lowercase" =>            ["lowercase"]
"Class" =>                ["Class"]
"MyClass" =>              ["My", "Class"]
"MyC" =>                  ["My", "C"]
"HTML" =>                 ["HTML"]
"PDFLoader" =>            ["PDF", "Loader"]
"AString" =>              ["A", "String"]
"SimpleXMLParser" =>      ["Simple", "XML", "Parser"]
"vimRPCPlugin" =>         ["vim", "RPC", "Plugin"]
"GL11Version" =>          ["GL", "11", "Version"]
"99Bottles" =>            ["99", "Bottles"]
"May5" =>                 ["May", "5"]
"BFG9000" =>              ["BFG", "9000"]
"BöseÜberraschung" =>     ["Böse", "Überraschung"]
"Two  spaces" =>          ["Two", "  ", "spaces"]
"BadUTF8\xe2\xe2\xa1" =>  ["BadUTF8\xe2\xe2\xa1"]
```
