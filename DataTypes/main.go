package main

import "fmt"
func main(){
	var userName = "John Doe"
	var age = 12
	var cGpa = 4.5
	var isAboard = true

	fmt.Println(userName)
	fmt.Println(age)
	fmt.Println(cGpa)
	fmt.Println(isAboard)

	// calculation with string
	var a = "AB"
	var b = "CD"
	var resultString = a + b
	fmt.Println(resultString)

	// calculation with numbers 
	var n1 = 5
	var n2 = 6
	var resultNumber = n1 + n2
	fmt.Println(resultNumber)


	// define type before assigning to value
	var collegeName  string
	var major string
	var totalSeat int

	collegeName = "Dr Maleka College"
	major = "Science"
	totalSeat = 300

	fmt.Println(collegeName,major,totalSeat)
}