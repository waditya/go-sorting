package main

import (
	"fmt"
	"sort"
)

func main() {
	vars := []int{5, 2, 0, 1, 7, 3, 6, 4, 11, 9}
	fmt.Println("Unsorted array : ", vars)
	sort.Ints(vars)
	fmt.Println("Sorted   array : ", vars)

	rawStrings := []string{"Learning", "Go", "using", "examples", "is", "fun!"}
	fmt.Println("String slice before sorting : ", rawStrings)
	sort.Strings(rawStrings)
	fmt.Println("String slice after sorting : ", rawStrings)
}
