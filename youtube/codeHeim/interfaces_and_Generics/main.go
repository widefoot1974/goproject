package main

import "fmt"

// Sortable is a constraint that requires the type
// to support comparison operators
type Sortable interface {
	~int | ~float64 | ~string
}

// Generic function to sort a slice of any Sortable type
func GenericSort[T Sortable](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	stringSlice := []string{
		"apple", "orange",
		"banana", "pear",
	}

	GenericSort(intSlice)
	GenericSort(stringSlice)

	fmt.Println("Sorted intSliace:", intSlice)
	fmt.Println("Sorted stringSlice:", stringSlice)
}
