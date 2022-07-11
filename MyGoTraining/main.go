package main

import "fmt"

func main() {
	// exercise 1 on slice (composite literal)
	// x := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(x)
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T", x)

	// // exercise 2 on slice
	// y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(y)
	// for i, v := range y {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T", y)

	// exercise 3 on map
	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])

}
