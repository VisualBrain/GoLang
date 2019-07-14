package main

import (
	"fmt"
	"strings"
)

/*
 To remove the error - exported type PlantStatus should have comment or be unexported
 write above the definition
// PlantStatus is ...
every type with fieldName - first letter capital - Go will identify it as exported component
e.g Agent
unexported struct --: agent
*/

func main() {
	//	plantCapacityPerHead := [...]float64{40., 50.3, 60., 18., 45, 70.8, 19.}
	//	plantCapacityPerHead := []float64{40., 50.3, 60., 18., 45, 70.8, 19.}

	plants1 := []PowerPlant{
		PowerPlant{wind, active, 100},
		PowerPlant{hydro, inactive, 250},
		PowerPlant{solar, unavailable, 150},
		PowerPlant{solar, active, 145},
		PowerPlant{hydro, unavailable, 150},
		PowerPlant{wind, active, 184},
	}
	//	activePlants := [...]int{0, 2}
	//activePlants := []int{0, 2}
	fmt.Println("\nHello Plant Management")
	fmt.Println("(1) Generate Power Plant Report \t", "(2) Generate Power Grid Report\n")
	println("Choose one of these options:")
	var choose int
	fmt.Scanln(&choose)
	grid := PowerGrid{300, plants1}
	switch choose {
	case 1:
		grid.generatePlantReport()
		//generatePowerPlantReport(plantCapacityPerHead...) //spread operator
	case 2:
		//	generatePowerGridReport(activePlants, plantCapacityPerHead, 75.0)
		grid.generateGridReport()
	default:
		println("Unknown Option choosen")
	}
}
func generatePowerPlantReport(plantCapacityPerHead ...float64) {
	for index, val := range plantCapacityPerHead {
		fmt.Printf("Plant No : %d has capacity : %.2f\n", index, val)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacityPerHead []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacityPerHead[plantID]
	}
	fmt.Printf("The Total Grid Capacity is: %.2f\n", capacity)
	fmt.Printf("The Load is : %.2f\n", gridLoad)
	fmt.Printf("The Utilization is : %.2f%%", gridLoad/capacity*100)
}

// PlantStatus is ...
type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Not_Available"
)

// PlantType is ...
type PlantType string

const (
	hydro PlantType = "Hydro_Electric_Power_Plant"
	wind  PlantType = "Wind_Electric_Power_Plant"
	solar PlantType = "Solar_Energy_Generator_Power_Plant"
)

// PowerPlant is ...
type PowerPlant struct {
	plantType   PlantType
	plantStatus PlantStatus
	capacity    float64
}

// PowerGrid is ...
type PowerGrid struct {
	loadGrid float64
	plants   []PowerPlant
}

func (powerGrid *PowerGrid) generatePlantReport() {
	for index, plant := range powerGrid.plants {
		label := fmt.Sprintf("%s%d", "The Plant index is :", index)
		fmt.Println(label)
		fmt.Println(strings.Repeat("_", len(label)))
		fmt.Printf("%s%s\n", "Plant Status is :", plant.plantStatus)
		fmt.Printf("%s%s\n", "Plant Type is :", plant.plantType)
		fmt.Printf("%s%f\n", "Plant Capacity is :", plant.capacity)
	}
}
func (powerGrid *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, plant := range powerGrid.plants {
		if plant.plantStatus == active {
			capacity += plant.capacity
		}
		//capacity += plantCapacityPerHead[plantID]
	}
	label := "Plant Grid Report "
	println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("The Total Grid Capacity is: %.2f\n", capacity)
	fmt.Printf("The Load is : %.2f\n", powerGrid.loadGrid)
	fmt.Printf("The Utilization is : %.2f%%", powerGrid.loadGrid/capacity*100)
}
