package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Error :
  idiometic to return an error as last returned value if any error occur
	func test(target typename)(result typename,err error)
	if err == nil {  // nil is used to indicate success
	  //no error
	}else{
	fmt.Println("Error occured:",err)
	}
}

*/
func main() {

	fmt.Println("Branch And Looping")

	/*
		// if initialization comparison_operations*/

	//	first way
	val := 1 //globally defined
	if val == 1 {
		println("True")
	} else {
		println("False")
	}

	//second way
	if val := 1; val > 0 { //local to if
		println("True")
	} else {
		println("False")
	}

	/*--------switch --------------------*/

	switch foo := 1; foo { //foo will be local to switch
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("Default")
	}
	switch random := randNum(); random {
	case 0, 2, 4, 6, 8:
		println("Odd Number")
	case 1, 3, 5, 7, 9:
		println("Even Number")
	default:
		println("Default")

	}
	foo := 1
	switch {
	case foo == 1:
		println("one")
		fallthrough
	case foo == 2:
		println("two")
	case foo >= 1:
		println("Greater than equal to 1")
	case foo > 2:
		println("Greater than 2")
	default:
		println("Default")
	}
	//third way
	switch foo {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("Default")
	}

	/*---------------loop -------------------*/
	//first way
	for i := 1; i < 5; i++ {
		fmt.Printf("%s%d\n", "Index is : ", i)
	}
	//second Way
	i := 1
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	array := [...]int{23, 13, 56, 23, 134, 59}
	for index, val := range array {
		fmt.Printf("%-5s%+20d%+11s%d\n", "The value of : ", index, " is : ", val)
	}

	Map := make(map[string]string)
	Map["One"] = "Happy Birthday"
	Map["two"] = "Good Night"
	Map["three"] = "Sweet Dream"

	for key, val := range Map {
		fmt.Printf("%-14s %-6s%-21s%s\n", "The Key is :", key, "And Value is :", val)
	}
}

/*
fallthrough()
still execute all other cases after found the correct one


*/

func randNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
