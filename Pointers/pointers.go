package main

import "fmt"

func main() {
	age := 32

	fmt.Println(&age)

	var agePointer *int = &age //declaring

	getAdultYearAge(agePointer)
	fmt.Println(age)
}

func getAdultYearAge(age *int) {
	fmt.Println(age) //geting same memory location
	*age = *age - 18 //dereferancing and getting actual value
}
