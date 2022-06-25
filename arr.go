package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr2d := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	arr3d := [1][2][3]int{{{1, 2, 3}, {1, 2, 3}}}
	arr4d := [1][2][2][3]int{{{{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}}}}
	arr5d := [1][2][2][2][3]int{{{{{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}}}, {{{1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}}}}}

	fmt.Println("1d: ", arr)
	fmt.Println("2d: ", arr2d)
	fmt.Println("3d: ", arr3d)
	fmt.Println("4d: ", arr4d)
	fmt.Println("5d: ", arr5d)
}
