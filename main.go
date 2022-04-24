package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	var bugPath string
	fmt.Print("Insert bug file path: ")
	_, err := fmt.Fscan(os.Stdin, &bugPath)
	check(err)

	var landscapePath string
	fmt.Print("Insert landscape file path: ")
	_, err = fmt.Fscan(os.Stdin, &landscapePath)
	check(err)

	bug, err := ioutil.ReadFile(bugPath)
	check(err)
	bugStr := spaceStringsBuilder(string(bug))

	landscape, err := ioutil.ReadFile(landscapePath)
	check(err)
	landscapeStr := spaceStringsBuilder(string(landscape))

	log.Println("Count:", strings.Count(landscapeStr, bugStr))
}

func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
