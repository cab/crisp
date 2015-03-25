package main

import "fmt"
import "github.com/cab/crisp"

func main() {
	var input = "(do \"this\")"
	fmt.Println(crisp.Read(input))
}
