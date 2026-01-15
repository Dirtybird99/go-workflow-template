package main

import (
	"flag"
	"fmt"

	"github.com/example/go-workflow-template/pkg/greet"
)

var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	fmt.Println(greet.Hello("World"))
}
