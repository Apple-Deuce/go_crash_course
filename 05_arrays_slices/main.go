package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and Assign
	//fruitArr := [2]string{"Apple", "Orange"}

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)

	//Amount of Values in Slice
	fmt.Println(len(fruitSlice))

	//Returning Partial Slice or Array
	fmt.Println(fruitSlice[1:3])
}
