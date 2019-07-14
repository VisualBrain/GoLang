package main

/*
only write the packages that you use ,importing unnecessary packages will show error
to bypass the error use _(underscore) in front of package then it will not show error whether it is used or not
e.g
   _ "os"

*/
import (
	"fmt"
	"os"
	_ "os" //bypassing standard go rule - don't import unnecessary packages
	"reflect"
	_ "reflect"
	"runtime"
	_ "runtime"
	"strings"
	_ "strings"
)

const (
	msg = " Hello I am Constant! constant - I cannot be changed"
)

//Package Level Variables are Global
var (
	msg1    = "Hello I am Variable! I will change according to surrounding"
	message string
	//var1, var2 string
	var1, var2 = "Variable 1", "Variable 2"
	var3       = 45.4
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
	var1 = "Helo from Variable 1"
	var2 = "Helo from Variable 2"
	var4 := 56 //shorthand declare only work inside function
	println(msg1, "\t", var1, "\t", var2, "\t", var4)
	fmt.Println("The type of Var1 is :", reflect.TypeOf(var1))
	println("Running Program Operating System:", strings.ToUpper(runtime.GOOS))
	fmt.Println("Profile Cpu :", runtime.CPUProfile)
	fmt.Println("Memory Profile", runtime.MemProfile)
	fmt.Println("Number of Cpu:", runtime.NumCPU)
	fmt.Println("The explicit typecast value of Var3 is:", int(var3))
	//fmt.Println("Environment Variables are :", os.Environ())
	for _, val := range os.Environ() {
		println(val)
	}
	println("User Name:", os.Getenv("USERNAME"))
}

//init - will run when its module is referenced by another module
//always init will run befor main function
func init() {
	msg1 = "Hello Init i will be d when the first time my name is being asked "
}

/*
About Pointer
&(a) gives the address of Variable - Reference a Pointer
b := &(a)
*(b) - De reference a Pointer
*(b) gives the value of variable present at address &(a)

*/
