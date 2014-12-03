# gistme

Gistme creates gists from command line.

## First things first

You'll need to install [go](http://golang.org/doc/install). You'll also 
need to to generate a personal [access token](https://github.com/settings/applications)

## Usage

### CLI

Creating a private gist

    go run examples/alpha.go -n "my name is Baby Doe" -d "bye bye"

Creating a public gist

    go run examples/alpha.go -p=1 -n "my name is Jane Doe" -d "Hello Hello"

Creating a gist with defaults. Note that the gist is always private unless
you pass in a flag p=1 or -privacy=1 or even --privacy=1. Also note that 
privacy is the only flag which demands an equal sign.

    go run examples/alpha.go

Piping some content

    cat snippet | go run examples/alpha.go
    cat README.md | go run examples/alpha.go -name="em branco" -d "story of my life" -p=1

Once you're satisfied with the example/alpha.go, you can build it once 
and use it forever

    go build examples/alpha.go

move the created binary file to `$HOME/go/bin` or anywhere under your `$PATH`.

For more info, please read the [![GoDoc](https://godoc.org/github.com/netp/gistme?status.svg)](https://godoc.org/github.com/netp/gistme)
[![Build Status](https://travis-ci.org/netp/gistme.svg?branch=master)](https://travis-ci.org/netp/gistme)


# With OSX and Vim

## Visual mode

    '<,'>!pbcopy && pbpaste | gistme
    u

## Insert mode

    :36;39!pbcopy && pbpaste | gistme
    u

## TODO

write an alias for `.vimrc`
