package main

import "fmt"

func main() {
	names := []string{"ali", "amir", "roola", "ebram"}
	fmt.Print("Enter a name to search: ")
	key := ""
	fmt.Scanln(&key)
	found := false
	index := -1

	for i, v := range names {
		if key == v {
			found = true
			index = i
			break
		}
	}
	if found {
		fmt.Printf("'%v' found at index: %v \n", key, index)
	} else {
		fmt.Printf("'%v' not in list \n", key)
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
