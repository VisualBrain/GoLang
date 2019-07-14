package main

import (
	"fmt"
)

func main() {

	fmt.Println("\nHello to Our Plant")
	var plantCapacityPerHead = []float64{}
	plantCapacityPerHead = []float64{40., 50.3, 60., 18., 45, 70.8, 19.}

	var totalCapacityOfPlant = plantCapacityPerHead[0] + plantCapacityPerHead[1] +
		plantCapacityPerHead[2] + plantCapacityPerHead[3] + plantCapacityPerHead[4] +
		plantCapacityPerHead[5] + plantCapacityPerHead[6]

	var gridload = 50.6
	var utitlization = totalCapacityOfPlant / gridload
	fmt.Printf("%-20s%.1f%%\n", "Total utitlization:", utitlization*100)         //right side of string space = 20 and precision = 1
	fmt.Printf("%-20s%.0f\n", "Total Capacity of Plant: ", totalCapacityOfPlant) //right side of string space = 20 and precision = 0
	fmt.Printf("%-20s%.0f\n", "Total Load: ", gridload)
}

/*
%-20s -- means space following string having 20 space provided for string dkkkd_______________
%s - string
%d - integer
%t - boolean
%x - hexadecimal -lowercase (a-f)
%X - hexadecimal - uppercase (A-F)
%o - octal
%v - universal
%p - pointer or slice - 16 bits
%% - %

%+20s - means 20 space-sizeofstring before string _______________dkkkd


*/
