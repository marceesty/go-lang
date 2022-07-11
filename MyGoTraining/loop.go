package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	for j := 65; j <= 90; j++ {
		fmt.Println(j)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	x := []int{2, 4, 6, 8}
	fmt.Println(x)
}
