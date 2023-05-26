package main

import (
	"flag"
	"fmt"
)

func defineFlags() (string, bool) {
	name := flag.String("name", "World", "Name to greet")
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	flag.Parse()
	return *name, *verbose
}

func greet(name string, verbose bool) {
	if verbose {
		fmt.Println("Running in verbose mode")
	}
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	name, verbose := defineFlags()
	greet(name, verbose)
}
