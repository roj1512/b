package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-]")

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(1)
	}
	fmt.Print(re.ReplaceAllString(strings.Join(args, " "), "-"))
}
