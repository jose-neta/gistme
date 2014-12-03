package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/netp/gistme"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	name        string
	is_private  bool
	description string
)

func init() {
	// long version
	flag.StringVar(&name, "name", "gist name", "name of the gist file")
	flag.BoolVar(&is_private, "private", false, "privacy")
	flag.StringVar(&description, "desc", "gist description", "gist description")

	// shorthand version
	flag.StringVar(&name, "n", "gist name", "name of the gist file")
	flag.BoolVar(&is_private, "p", false, "privacy")
	flag.StringVar(&description, "d", "gist description", "description of gist")
}

// main is a command line utility which allow you to create gists
func main() {
	flag.Parse()

  if terminal.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Println("Type your message, enter ^D to send, ^C to abort:")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	g := gistme.NewGist()

	// CREATE A NEW GIST
	res := g.Create(name, description, is_private, strings.Join(lines, "\n"))
	fmt.Println(res)

	// GET ALL MY GISTS
	// fmt.Println(g.All())

}
