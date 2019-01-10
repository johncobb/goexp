package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("goexp protobuf example:")

	vehicle := &Vehicle {
		Vin: "WP1AA2AY0KDA00835",
		Year: 2019,
		Make: "Porsche",
		Model: "Cayenne",
		Color: "Blue",
		Type: Vehicle_NEW,
		VehicleFinancial: &VehicleFinancial {
			Price: 80000,
			Cost: 70000,
		},

	}

	data, err := proto.Marshal(vehicle)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}
	fmt.Println(data)

	newVehicle := &Vehicle{}
	err = proto.Unmarshal(data, newVehicle)
	if err !=nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Println(newVehicle.GetVin())
	fmt.Println(newVehicle.GetYear())
	fmt.Println(newVehicle.GetMake())
	fmt.Println(newVehicle.GetModel())
	fmt.Println(newVehicle.GetColor())
	fmt.Println(newVehicle.GetType())
	fmt.Println(newVehicle.VehicleFinancial.GetPrice())
	fmt.Println(newVehicle.VehicleFinancial.GetCost())


}
