# gistme [![GoDoc](https://godoc.org/github.com/netp/gistme?status.svg)](https://godoc.org/github.com/netp/gistme) [![Build Status](https://travis-ci.org/netp/gistme.svg?branch=master)](https://travis-ci.org/netp/gistme)


Gistme creates gists from CLI.

## First things first

You'll need to install [go](http://golang.org/doc/install) and you'll also 
need to to generate a personal [access token](https://github.com/settings/applications)

## Usage

### CLI

Note that every following command assume that you've the value of your 
token passed via environment variable or by flag -token (-t). 

If you hit `gistme` alone you'll get

    gistme

      You must define the token value via env var $TOKEN
      OR pass the token value via -token (-t) flag.
      Please see gistme -h

You must pass the token either by flat -token (-t)

    gistme -t `cat sandbox/token`

      Type your message, enter ^D to send, ^C to abort:
      done
      https://gist.github.com/1917fbcbe922c4d290af

or by environment variable

    TOKEN=`cat sandbox/token` gistme

      Type your message, enter ^D to send, ^C to abort:
      env var way

Fell free to check the help

    gistme -h

      Usage of gistme:
        -d="gist description": description of gist
        -desc="gist description": gist description
        -n="gist name": name of the gist file
        -name="gist name": name of the gist file
        -p=false: privacy
        -private=false: privacy
        -t="": your user token
        -token="": your user token

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

You can modify the `example/alpha.go` in order to suits you best. Once 
you're satisfied with mofications, build it once and use it forever.

    go build examples/alpha.go

Now move the created binary file to `$HOME/go/bin` or anywhere under 
your `$PATH`.

### With Vim on OSX

#### Visual mode

    '<,'>!pbcopy && pbpaste | gistme
    u

#### Insert mode

    :36;39!pbcopy && pbpaste | gistme
    u

# TODO

Write an alias for `.vimrc`
