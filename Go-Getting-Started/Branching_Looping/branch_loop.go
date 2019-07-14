package main

import (
	"fmt"
)

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

	foo := 1
	switch {
	case foo == 1:
		println("one")
	case foo == 2:
		println("two")
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
