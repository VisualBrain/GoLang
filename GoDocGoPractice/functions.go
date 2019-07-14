//Variadic Functions,Recursive Functions,Annonymous Functions
package main

import (
    "fmt"
    )

//Variadic Functions
func sumIn(numbers ...int) int { // take any number of arguments
    total := 0
    for _,num := range numbers {
        total+=num
    }
    return total
}
//Annonymous functions 
func annoymousFunctions() func() int {  //This function annoymousFunctions returns another function
    i:=0
    return func() int {
        i=i+1
        return i
    }
}
func recursiveFactorial(Number int) int {
    if Number == 0 {
        return 1
    }
    return Number * recursiveFactorial(Number-1)
}
func main() {
    var sum int = sumIn(1,3,4,5,7,8,4,2,46)
    fmt.Println(sum)
    nextnum := annoymousFunctions()
    fmt.Println(nextnum())
    fmt.Println(nextnum())
    fmt.Println(nextnum())
    fmt.Println(nextnum())
    fmt.Println(recursiveFactorial(5))
}
