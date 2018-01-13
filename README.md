go-gemini
==========

go-gemini is an implementation of the Gemini API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
	import "github.com/jyap808/go-gemini"

This library is more of a framework for some bots I use so it is expected that a lot of things don't work but pull requests are accepted.

## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/jyap808/go-gemini"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Gemini client
	gemini := gemini.New(API_KEY, API_SECRET)

	// Get markets
	markets, err := gemini.GetMarkets()
	fmt.Println(err, markets)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-gemini/blob/master/examples/gemini.go)

