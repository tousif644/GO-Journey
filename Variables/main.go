package main

import "fmt"

func main(){
	var conferenceName = "Go conference"
	fmt.Println(conferenceName)

	var a = "apple"
	fmt.Println(a)

	var b,c = 1, 2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int = 15
	fmt.Println(e)

	// const variables
	const price = 100
	fmt.Println(price)
	// price = 200 
	// fmt.Println(price)


	// var is changeable const is not changeable
}