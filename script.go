// package main

import "fmt"

// start program
func main() {
	var name string = "Abdelilah"

	age, likesGo := 20, true
	// fmt.Println("Hello World!")
	fmt.Println("Hi " + name)
	fmt.Println("Your age: ", age)
	fmt.Println("You like Golang: ", likesGo)

	// data structure: Arrays
	var myArray [3]string
	myArray[0] = "😇"
	myArray[1] = "🙂"
	myArray[2] = "🤭"
	fmt.Println(myArray)

	// Maps 
	var myMap = make(map[string]string)
	myMap["robot"] = "🤖"
	myMap["monster"] = "👾"
	myMap["alien"] = "🤡"
	fmt.Println(myMap)


	// Control Flow
	animal  := "🐶"
	if animal == "🐶" {
		fmt.Println("🙋")
	} else {
		fmt.Println("😞")
	}
	

	// loops
	for i := 0; i < 10; i++ {
		fmt.Println("🤠")
	}


	// pointers
	var year int = 2023
	var p *int = &year
	fmt.Println(p)

}
