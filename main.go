package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Missing Your Name :(")
		os.Exit(1)
	}

	name := args[0]
	if name == "" {
		fmt.Printf("Empty Name :(")
		os.Exit(1)
	}

	fmt.Printf("Hello %s\n", name)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hello [yourname]\n")
	flag.PrintDefaults()
	os.Exit(2)
}
