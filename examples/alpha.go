package main

import (
	"bufio"
  "log"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/netp/gistme"
	"golang.org/x/crypto/ssh/terminal"
)

const ( usage_msg = `

You must define the token value via env var $TOKEN
OR pass the token value via -token (-t) flag. 
Please see gistme -h

`)

var (
	name        string
	is_private  bool
	description string
  token       = os.Getenv("TOKEN")
)

func init() {

	// long version
	flag.StringVar(&name, "name", "gist name", "name of the gist file")
	flag.BoolVar(&is_private, "private", false, "privacy")
	flag.StringVar(&description, "desc", "gist description", "gist description")
	flag.StringVar(&token, "token", token , "your user token")

	// shorthand version
	flag.StringVar(&name, "n", "gist name", "name of the gist file")
	flag.BoolVar(&is_private, "p", false, "privacy")
	flag.StringVar(&description, "d", "gist description", "description of gist")
	flag.StringVar(&token, "t", token , "your user token")
}

// main is a command line utility which allow you to create gists
func main() {
	flag.Parse()
  if token == "" {
    log.Fatal(usage_msg)
  }

  if terminal.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Println("Type your message, enter ^D to send, ^C to abort:")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	g := gistme.NewGist( token )

	// CREATE A NEW GIST
	res := g.Create(name, description, is_private, strings.Join(lines, "\n"))
	fmt.Println(res)

	// GET ALL MY GISTS
	// fmt.Println(g.All())

}
