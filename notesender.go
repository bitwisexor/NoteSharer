package ns

import (
	"flag"
	"fmt"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "file name")
	flag.StringVar(&filename, "f", "", "file name (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -f {filename}\n", os.Args[0])
	}
}

func main() {
	flag.Parse()

	if filename == "" {
		fmt.Println("Error: Filename required.")
		flag.Usage()
		os.Exit(1)
	}
}
