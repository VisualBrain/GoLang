package main

import (
	"fmt"
)

var slice []string

func main() {
	fmt.Println("\nFunctions and Parameters")
	sayJekyll("Verbise")
	sayJekyll("Angular")
	sayJekyll("Go")
	fmt.Println(slice)

	messRef := "The message is en"
	sayJekyllByReference(&messRef)
	println(messRef)
	variadicFunctions("ksksksk", "Gandpa", "Verilog", "Microservices", "Jumsi Tana", "Pascal", "COBOL")
	list, length := addNames("Jar", "War", "Avi", "Mp4", "Mp3", "MPEG", "MPEG4", "MKV")
	fmt.Println(list, "\t Length is:", length)
	list1, length1 := addNamedReturnValues("Jar", "War", "Avi", "Mp4", "Mp3", "MPEG", "MPEG4", "MKV")
	fmt.Println(" By Named :", list1, "\t Length is:", length1)

	//Annonymous Function
	addNamesUsingAnnonymousFunctions := func(names ...string) ([]string, int) {
		list1 := make([]string, len(names))
		for _, val := range names {
			list1 = append(list1, val)
		}
		length1 = len(list1)
		return list1, length1
	}
	list2, length2 := addNamesUsingAnnonymousFunctions("ksksksk", "Gandpa", "Verilog", "Microservices", "Jumsi Tana", "Pascal", "COBOL")
	fmt.Println(" By Annonymous Functions :", list2, "\t Length is:", length2)

}
func sayJekyll(message string) { //by value take copy of original value and do operations on it without changing the original value
	slice = append(slice, message)
	println("Message Received :", message)
}
func sayJekyllByReference(message *string) { //By reference change the orignal value
	println(*message)
	*message = "Hello JeyKll by reference"
}
func variadicFunctions(val ...string) { //list of values       Variadic Function
	fmt.Println(val)
}
func addNames(names ...string) ([]string, int) {
	memory := make([]string, len(names))
	for _, val := range names {
		memory = append(memory, val)
	}
	return memory, len(memory)
}

//named Return values
func addNamedReturnValues(names ...string) (list1 []string, length1 int) { // return names same as return names in main-function
	list1 = make([]string, len(names))
	for _, val := range names {
		list1 = append(list1, val)
	}
	length1 = len(list1)
	return list1, length1
}
