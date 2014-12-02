# gistme

Gistme creates gists from command line.

## First thing first

You'll need to install [http://golang.org/doc/install](go). You'll also 
need to to generate a personal [access token](https://github.com/settings/applications)

## Usage

### cli side

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


# TODO 

- show usage inside vim 
