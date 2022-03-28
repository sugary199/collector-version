// main.go

package main

import (
	"flag"
	"fmt"


	"github.com/sugary199/collector-version/version"
)

func main() {
	var ver = flag.Bool("version", false, "show version")
	flag.Parse()

	if *ver {
		fmt.Printf("%s\n", version.Version())
		return
	}
}