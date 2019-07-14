package main

import (
	"fmt"
)

//struct is container that hold arbitary data types called field
/*
In Go ,there is no class or constructor concepts, instead we used function to behave like constructor
GO definition of constructor : a function whose only purpose to create properly initialize object and pass it to the caller

*/
type structName struct {
	nameOf string
}

type construct struct {
	Map map[string]string
}
type messagePrinter struct {
	message string
}

func main() {
	visa := structName{"Visha"}
	//  println(visa)  -error - illegal types for operand:print structName
	//to get the Address
	visa3 := &structName{"Vishae"}
	println("\nVisa 3 address is:", visa3)
	fmt.Println(visa)
	visa1 := new(structName)
	visa1.nameOf = "Aditya"
	println("Address of visa 1 is:", visa1)
	println(visa1.nameOf)
	fmt.Println("value is:", visa.nameOf)

	/*
		createConstruct := &construct{}
		createConstruct.Map["bare"] = "cover"
		println(createConstruct) //Error : assignment to entry in nil map
	*/
	createConstruct := newCreateConstructor()
	createConstruct.Map["bare"] = "cover"
	fmt.Println(createConstruct)

	//Inheriting basefunctionaliy from parent to child
	mess := messagePrinter{"A message is Received"}
	mess.printMessage() //using method
	messInheritance := enhancedPrinter{}
	messInheritance.message = " Go used composition"
	//second way
	messInheritance1 := enhancedPrinter{messagePrinter{"Message"}}
	messInheritance.printMessage()
	messInheritance1.printMessage()
}

//creating constructor
func newCreateConstructor() *construct { //mirror of constructor in c++ class
	result := construct{}
	result.Map = map[string]string{}
	return &result
}

/*
A method is a special function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
A "Receiver" is present to represent the container of the function
This receiver can be used to call a function using "." operator
e.g
   func (receiverName receiverType) functionName() return_type {
   // Write something here
 }
*/
func (messagePrint *messagePrinter) printMessage() { //Struct Methods -  A mirror of methods in c++ class
	println(messagePrint.message)
}

type enhancedPrinter struct {
	messagePrinter
}
