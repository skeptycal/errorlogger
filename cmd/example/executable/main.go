package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/errorlogger"
	. "github.com/skeptycal/errorlogger/cmd/example/executable/osargsutils"
)

var log = errorlogger.Log

func main() {
	fmt.Printf("%25.25s %s\n", "raw os.Args[0]:", os.Args[0])

	arg0, err := Arg0()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%25.25s %s\n", "using Arg0():", arg0)

}
