// main.go

package main

import (
	"flag"
	"fmt"

	"collecotr-version/core"
)

func main() {
	var ver = flag.Bool("version", false, "show version")
	flag.Parse()

	if *ver {
		fmt.Printf("%s\n", core.Version())
		return
	}
}