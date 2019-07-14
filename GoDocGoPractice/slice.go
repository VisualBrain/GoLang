//Slices
/*
The type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.
*/
package main

import (
    "fmt"
    "math/rand"
    "sort"
)

func main() {

    /*
    A slice can be created with the built-in function called make, which has the signature,

    func make([]T, len, cap) []T
    */
    //First Way -Different - we need not to put the number of element in []
    arrayTypeSlice := []int {1,3,45,6,7,4,768,44}
    fmt.Println(arrayTypeSlice)
    //Second Way
    makeSlice := make([]int , 6,6)  //len = 0, capacity =6
    fmt.Println(makeSlice) //Print  Empty Slice
    makeSlice[3] = 383
    fmt.Println("The Slice is:",makeSlice) //Does not append as The len of makeSlice is 0 the alternate way is using append
    makeSlice = append(makeSlice,388,44,22)
    fmt.Println(makeSlice,len(makeSlice),cap(makeSlice))
    capt()
    newSliceAfterRemoval := removeAtIndexSlice(makeSlice,3)
    fmt.Println("New Slice is:",newSliceAfterRemoval)
    
    twoDimension := make([][]int,4)
    for i:=0;i<4 ; i++ {
        innerLength := i+1
        twoDimension[i] = make([]int ,innerLength)
        for j:=0;j<innerLength;j++ {
            twoDimension[i][j] = i+j
        }
    }
    fmt.Println(twoDimension)
    RandomNizationSlice()
    
}

func capt() {
    fmt.Println("Executing External Function")
    scores := make([]int ,0,6)
    capacity:= cap(scores)
    fmt.Println("The initial Capacity is:",capacity,"And Slice score is :",scores)
    for i:=0;i<25;i++ {
        scores = append(scores,i*5)
    }
    capacity= cap(scores)

    fmt.Println("After Append----------------------------------------------- \nScores is:",scores,"\n After appending Capacity is:",capacity)
}
func removeAtIndexSlice(tomArray []int,pos int) []int {
    /*
    //First way
    copy(tomArray[pos:],tomArray[pos+1:])
    */
    
    /* 
    //Second Way
      tomArray=append(tomArray[:pos],tomArray[pos+1:]...)
      return tomArray
    */
    
    
   /* 
   //Third way if order is not Important
      tomArray[pos],tomArray[len(tomArray)-1] = tomArray[len(tomArray)-1],tomArray[pos]
      return tomArray[:len(tomArray)-1]
   */
   //Fourth Way - it also don't follow order
    tomArray[pos] = tomArray[len(tomArray)-1]
    return tomArray[:len(tomArray)-1]
}

func RandomNizationSlice() {
    fmt.Println("Random Array Slice initialization")
    scores := make([]int,100)
    for i:=0;i<100;i++{
        scores[i] = int(rand.Int31n(1000))
    }
    fmt.Println(scores)
    sort.Ints(scores)
    fmt.Println(scores)

}

