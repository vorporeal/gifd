// gifd.go
package main

import (
	"fmt"
)

func main() {
	flags.Init()
	fmt.Println(flags.Port)
	fmt.Println(flags.Path)
}
