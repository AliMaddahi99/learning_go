package main

import (
	"fmt"
	"strings"
	// "sort"
)

func main() {
	arr := usersName()
	// sort.Strings(arr)
	// fmt.Println(arr)

	sortUsersName(arr)
	fmt.Println(arr)

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}

func usersName() []string {
	fmt.Println("Enter Users name, separate them with comma: ")
	userNames := ""
	fmt.Scanln(&userNames)
	return strings.Split(userNames, ",")
}

func sortUsersName(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}
