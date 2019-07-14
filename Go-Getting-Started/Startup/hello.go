package main

/*
only write the packages that you use ,importing unnecessary packages will show error
to bypass the error use _(underscore) in front of package then it will not show error whether it is used or not
e.g
   _ "os"

*/
import (
	"fmt"
	_ "os" //bypassing standard go rule - don't import unnecessary packages
)

const (
	msg = " Hello I am Constant! constant - I cannot be changed"
)

var (
	msg1    = "Hello I am Variable! I will change according to surrounding"
	message string
)

/*
default values-
  string = ''
  variable = 0
*/
/*
  1. using fmt
  2. using var,const
  3. using

*/
func main() {
	fmt.Println("Hello World")
	//	msg = "Hello" //show Error cannot assign to msg
	println(msg) //second Way without using fmts
	println(msg1)
	println("My Message :", message)
	msg1 = "Hello I am Variable! I can be change"
	println(msg1)

}

//init - will run when its module is referenced by another module
//always init will run befor main function
func init() {
	msg1 = "Hello Init i will be d when the first time my name is being asked "
}
