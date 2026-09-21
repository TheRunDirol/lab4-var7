package main

import (
	"fmt"

	"lab4-var7/pkg/fueltrip"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	var distance float64
	var consumption float64
	var price float64
	var reserve float64
	var route string

	fmt.Print("Введите маршрут: ")
	fmt.Scanln(&route)

	fmt.Print("Введите расстояние в км: ")
	fmt.Scan(&distance)

	fmt.Print("Введите расход топлива на 100 км: ")
	fmt.Scan(&consumption)

	fmt.Print("Введите цену топлива за литр: ")
	fmt.Scan(&price)

	fmt.Print("Введите запас топлива в процентах: ")
	fmt.Scan(&reserve)

	fuel, err := fueltrip.FuelNeeded(distance, consumption)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	cost, err := fueltrip.TripCost(fuel, price)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = fueltrip.AddTrafficReserve(&fuel, reserve)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	report, err := fueltrip.FormatTripReport(
		route,
		fuel,
		cost,
	)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	operationID := uuid.New().String()

	color.Cyan("Номер операции: %s", operationID)

	fmt.Printf("\nНеобходимое топливо: %.2f л\n", fuel)
	fmt.Printf("Стоимость поездки: %.2f тг\n", cost)
	fmt.Printf("%s\n", report)
}
