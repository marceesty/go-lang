package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("\n%d\t%b\t%U\t%#x\t%x\t%X", i, i, i, i, i, i)
	}
}
