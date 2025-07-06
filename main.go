package main

import "fmt"

func main() {
	arrayOfNames := [5]string{"a", "b", "c", "d", "e"} // array
	fmt.Println(arrayOfNames)

	sliceOfNames := []string{"a","b","c"} // slice
	sliceOfNames = append(sliceOfNames,"d") // Append returns the updated slice.
	fmt.Println(sliceOfNames)
}
