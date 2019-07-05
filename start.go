package main

/*
   if variable is new than use :=
   otherwise to update the previous declared variable use =
   function can return mutliple values
*/
import (
	"fmt"
	"os"
)

func main() {
	var name string
	//var age int
	//age =45
	//other way age:=getAge()
	//shorter way
	education := "BSLC"
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Time is Over", os.Args[1])
	fmt.Println("Hello, playground", education, name, getAge())
	sof, ware := getMessage("boolean")
	fmt.Println("vALUE OF sOFT", sof, ware)
	_, was := getMessage("Only get one")
	fmt.Println("vALUE OF sOFT", was)
	pre, _ := getMessage("Only get one")
	fmt.Println("vALUE OF sOFT", pre)

}
func getAge() int {
	return 45
}
func getMessage(arfu string) (string, bool) {
	if len(arfu) >= 10 {
		return arfu, false
	} else {
		return arfu, true
	}
}
