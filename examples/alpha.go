package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/netp/gistme"
)

func main() {
	// FIXME -- 2014-11-29 14:03 UTC --JPN --
	// Print this line only if there's no input file comming from
	// STDIN
	fmt.Println("Type your message, enter ^D to send, ^C to abort:")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	g := gistme.NewGist()

	// CREATE A NEW GIST
	res := g.Create("gist name", "gist description", true, strings.Join(lines, "\n"))
	fmt.Println(res)

	// GET ALL MY GISTS
	// fmt.Println(g.All())
}
