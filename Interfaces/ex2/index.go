package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct {
	brand string
	color string
}

type Bike struct {
	brand string
	color string
}

func (c Car) Drive() string {
	return "Driving a " + c.brand + " " + c.color + " car"
}

func (b Bike) Drive() string {
	return "Riding a " + b.brand + " " + b.color + " bike"
}

func getDrivingVehicleInfo(vehicle Vehicle) {
	fmt.Println(vehicle.Drive())
}
func main() {
	//new example of interface
	//create a new car
	car1 := Car{
		brand: "Toyota",
		color: "Red",
	}

	car2 := Car{
		brand: "Hyundai",
		color: "Blue",
	}

	bike1 := Bike{
		brand: "Honda",
		color: "Green",
	}

	bike2 := Bike{
		brand: "Royal Enfeild",
		color: "Red",
	}
	getDrivingVehicleInfo(car1)
	getDrivingVehicleInfo(car2)
	getDrivingVehicleInfo(bike1)
	getDrivingVehicleInfo(bike2)
}
