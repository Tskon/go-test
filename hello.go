package main

import (
	"fmt"
	"utils"
)

func main() {
	slice := make([]int, 4, 7)
	slice2 := []int{2, 0, 5, 4}
	slice2 = append(slice2, 3, 2)

	fmt.Println("hello world")
	fmt.Println(utils.Sum(1, 2))
	fmt.Println(utils.Sum2)
	fmt.Println(utils.A)
	fmt.Println(slice, cap(slice2))
}
