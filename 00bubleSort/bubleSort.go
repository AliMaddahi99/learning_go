package main

import "fmt"

// var age int = 22
// var name string = "Ali"
// var male bool = true
// var weight float32 = 65.5
// var arr = [10]int{0, 6, 8, 5, 3, 9, 1, 2, 4, 7}
// var slice = []int{2, 3, 0, 1, 4, 5}

func main() {
	length := 0
	fmt.Print("Enter the number of inputs: ")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs: ")
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&slice[i])
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = temp
			}
		}
	}

	fmt.Println(slice)
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
