package main

import (
    "fmt"
    )
 type Person struct {
         Name string
         Profile map[string]*Person
     }
func main() {
        newMap := make(map[string]int)
        newMap["A"]=1
        newMap["B"]=2
        newMap["C"]=3
        fmt.Println(newMap,len(newMap))
        power,exists := newMap["C"]  //3 true
        fmt.Println(power,exists)
        delete(newMap,"C")
        fmt.Println(newMap,len(newMap))
        
        //we can set the initialSize of Map
     //   map2 := make(map[string]int,100) //100 is the initialSize(how many keys will be in map)
     
     /*---------------------Map in Structure------------------------*/
    
     cod:= utilizePerson()
     fmt.Println(cod)
     cod.Profile["Visia"] = nil
     fmt.Println(cod)

}
func utilizePerson() *Person{
    return &Person{
        Name: "visva",
        Profile: make(map[string]*Person),
    }
}
