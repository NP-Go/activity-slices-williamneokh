package main

import "fmt"

func main() {
	//Insert Code here
	// Activity #1 Declare a new slice
	spent := []float64{9.5, 8, 10.2, 7.4}
	fmt.Printf("The lenght of the  slice is: %v\nThe capacity of the slice is: %v\n", len(spent), cap(spent))

	//Activity #2 Access the element of the slice
	fmt.Printf("Week 3 spending is: %.2f\n", spent[2])
	spent[2] = 9.8
	fmt.Printf("After amending week 3 spending is now: %.2f\n", spent[2])

	//Activity #3 Capacity after append
	spent = append(spent, 8.4, 9.4, 7.2)
	fmt.Printf("The lenght of the  slice is: %v\nThe capacity of the slice is: %v\n", len(spent), cap(spent))

	//Activity #4 Slicing
	var newSpent []float64
	newSpent = spent[3:]

	fmt.Printf("The lenght of the new slice is: %v\nThe capacity of the new slice is: %v\n", len(newSpent), cap(newSpent))
}
