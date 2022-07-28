package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// defining array
	var names [5]string
	names[0] = "Tousif"
	names[1] = "Tasrik"
	names[2] = "John"
	names[3] = "Doe"

	fmt.Println(names)

	var rolls [4]int
	rolls[0] = 3
	rolls[1] = 4
	rolls[2] = 5
	rolls[3] = 1

	fmt.Println(rolls)
	fmt.Println(rolls[2])

	//check the ype of array
	fmt.Printf("The type of this array is %T\n", names)

	//check the length of this array
	fmt.Printf("The length of this array is : %d\n", len(rolls))
	fmt.Printf("The length of this array is : %d\n", len(names))

	//slice the array
	sliceArray := []string{}

	// appending into sliceArray
	sliceArray = append(sliceArray, "potato", "raddish", "cauliflower", "cabbage")
	fmt.Println(sliceArray)

	// getting value by index
	fmt.Printf("%v\n", sliceArray[2])

	// Most easy way to append items into array by slice

	// appending sliceArray2 by int
	sliceArray2 := []int{}
	sliceArray2 = append(sliceArray2, 3, 4, 5, 6, 8, 9)
	fmt.Printf("The value of sliceArray2 is %v\n", sliceArray2)

	//concat array
	concated := fmt.Sprint(sliceArray2, sliceArray)
	// concated := fmt.Sprintf(sliceArray2, sliceArray)
	fmt.Printf("%v\n", concated)
}
