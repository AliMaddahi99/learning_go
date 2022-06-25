package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"ali", "xavi", "amir", "messi", "roola", "ebram", "ronaldo"}
	first := 0
	last := len(names)
	middle := (first + last) / 2
	found := false

	sort.Strings(names)

	fmt.Print("Enter a name to search: ")
	key := ""
	fmt.Scanln(&key)

	for last >= first {
		middle = (first + last) / 2
		if names[middle] == key {
			found = true
			break
		} else if names[middle] > key {
			last = middle - 1
		} else {
			first = middle + 1
		}
	}
	if found {
		fmt.Printf("%v found at index: %v \n", key, middle)
	} else {
		fmt.Printf("%v not found! \n", key)
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
