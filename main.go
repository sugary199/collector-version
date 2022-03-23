// main.go

package main

import (
	"flag"
	"fmt"


	"github.com/sugary199/collector-version/core"
)

func main() {
	var ver = flag.Bool("version", false, "show version")
	flag.Parse()

	if *ver {
		fmt.Printf("%s\n", core.Version())
		return
	}
}