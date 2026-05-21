package main

import (
	"fmt"

	"github.com/kuche1/imotbg/extract"
)

func main() {
	houses := extract.Main()

	for house := range houses {
		repr := house.Sprintf()
		fmt.Printf("%v\n", repr)
	}
}
