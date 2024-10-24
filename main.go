package main

import (
	"fmt"
	"os"
)

func main() {
	cmds := os.Args[1:]
	for _, cmd := range cmds {
		fmt.Printf("[%s]\n", cmd)
	}
}
