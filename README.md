# Galapagos

## Initial Setup

### Setup #1: Repository

1. Place the `galapagos` folder in your `GOCODE` src folder, usually `~/code/gocode/src`.

### Setup #2: Go Tools

**Note: Don't install Go using Homebrew, or autocomplete will not work in Atom correctly.**

1. Install Go as a binary distribution from: https://golang.org/

2. Install Godep (for dependency management):

       $ go get github.com/tools/godep

### Setup #3: Setup your editor

Atom is the recommended text editor for Go source.

1. Install Atom.

2. Install the following packages in Atom:
 * autocomplete-plus
 * go-plus

## Development

### Adding a new dependency

**Note: The only time you need to use godep is when you add a new dependency to an application or update an existing dependency that is already vendored in your application.**

1. Use go-get to get the dependency (EXAMPLE):

       $ go get -u github.com/russross/blackfriday

2. Add a relevant import in a Go file:

       import "github.com/russross/blackfriday"

3. Vendor the additional import (records it, copies it to Godeps/, rewrites imports):

       $ godep save -r ./...

### Updating an existing dependency

    $ go get -u <dependency>

    $ godep update <dependency>
