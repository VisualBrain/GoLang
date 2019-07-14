package main

import "fmt"

/*
Slice:
 Slice is an subset of Array,Slice is more flexible than Array,
the advantage over underlying array is : we can add many element as we want in slice and go will take care of resizing the underlying array
*/
/*  iota gives autoincrement feature*/
const (
	first = "First" //if we don't provide the values then it will copy the previous const assigned value to next const
	second
	third
	fourth = iota
	fifth
	sixth
)

func main() {

	/*------------------------------Data Types -----------------------*/
	// var myInt int
	// myInt = 43
	// println(myInt)
	// var myFloat float32 = 42.
	// println(myFloat)
	// fmt.Println("Hello")
	// myString := "Hello String"
	// println(myString)
	// myComplex := complex(4.5, 5)
	// var myComplex12 complex64
	// myComplex12 = complex(4.6, 9)
	// println(imag(myComplex12))
	// println(real(myComplex12))
	// println(myComplex, myComplex12)

	/*---------------------Constants------------*/

	println(first)  //first
	println(second) //first
	println(third)  //first
	println(fourth) //3
	println(fifth)  //4
	println(sixth)  //5

	myArray := [4]int{}
	//second way myArray := [4]int {345,1234,23,54}
	//the advantage of array in go - don't need to tell the size of array explicitly,compiler will done that job
	//e.g  myArray := [...]int {34,12,66,33}  //variadic array - we can infinite number of elements in this array - for every new addition - it creates new array
	//new_array = copy_of_previous_array+ new_element
	println("Value of Array before Initilization :", myArray[0])
	myArray[0] = 554
	myArray[1] = 552
	myArray[2] = 532
	myArray[3] = 434
	fmt.Println(myArray, len(myArray))
	mySlice := myArray[:]   //CREATING SLICE OF ALL THE VALUE IN myArray [554,552,532,434]  array[:] --> array[0] to array[n]
	mySlice1 := myArray[:3] //CREATING SLICE OF ARRAY FROM myArray[0] TO myArray[n-1]  [554,552,532]  array[:n] --> array[0] to array[n-1]
	mySlice2 := myArray[1:] //CREATING SLICE FROM myArray[1] EXCLUDING myArray[1]  array[s:] --> array[s+1] to array[n]
	fmt.Println(mySlice, mySlice1, mySlice2)
	//adding element to mySlice
	mySlice = append(mySlice, 449)
	fmt.Println("After Addition to Slice :", mySlice) //[554,552,532,434,449]

	sliceWithoutUnderlyingArray := []int{44, 23, 13, 64, 42} //now, we don't need to provide the size
	fmt.Println(sliceWithoutUnderlyingArray)

	//to solve the incompetency of sliceWithoutUnderlyingArray to maintain large data Array  we use make function
	// make(type, initialsize,capacity)
	slicewithMake := make([]float64, 100)
	slicewithMake[0] = 45
	slicewithMake[1] = 54
	slicewithMake[2] = 99

	fmt.Println(slicewithMake)
	//	slicewithMake1 := make([]float64, 100, 50) //len larger than capacity of slice
	slicewithMake1 := make([]float64, 100, 150) //len larger than capacity of slice

	slicewithMake1[0] = 45
	slicewithMake1[1] = 54
	slicewithMake1[2] = 199
	fmt.Println(slicewithMake1)

	/* the result to be
	   [45 54 99 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	    0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	*/
	/*-------------using map ------------------*/

	sliceMap := make(map[int]string) //map[key] valueType
	println(sliceMap)                //address of sliceMap
	fmt.Println(sliceMap)            // map[]
	sliceMap[45] = "439"
	sliceMap[34] = "4994"
	fmt.Println(sliceMap) // map[45:439 34:4994]

	postIncrement := 1
	postIncrement++
	println(postIncrement)
	postIncrement += 4 //Augmented Addition
	println(postIncrement)
	postIncrement *= 4
	println(postIncrement)
	postIncrement -= 13
	println(postIncrement)
	postIncrement %= 3 //Remainder
	println(postIncrement)

}
