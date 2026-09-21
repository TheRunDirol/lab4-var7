// Package fueltrip содержит функции для расчета поездки.
package fueltrip

import "fmt"

// FuelNeeded рассчитывает необходимое количество топлива для поездки.
func FuelNeeded(distanceKm, litersPer100 float64) (float64, error) {
	if distanceKm <= 0 {
		return 0, fmt.Errorf("расстояние должно быть больше нуля")
	}

	if litersPer100 <= 0 {
		return 0, fmt.Errorf("расход топлива должен быть больше нуля")
	}

	return distanceKm * litersPer100 / 100, nil
}

// TripCost рассчитывает стоимость топлива для поездки.
func TripCost(fuelLiters, pricePerLiter float64) (float64, error) {
	if fuelLiters < 0 {
		return 0, fmt.Errorf("количество топлива не может быть отрицательным")
	}

	if pricePerLiter <= 0 {
		return 0, fmt.Errorf("цена топлива должна быть больше нуля")
	}

	return fuelLiters * pricePerLiter, nil
}

// AddTrafficReserve добавляет запас топлива на случай пробок.
func AddTrafficReserve(fuel *float64, percent float64) error {
	if fuel == nil {
		return fmt.Errorf("указатель на топливо равен nil")
	}

	if *fuel < 0 {
		return fmt.Errorf("количество топлива не может быть отрицательным")
	}

	if percent < 0 || percent > 100 {
		return fmt.Errorf("процент запаса должен быть от 0 до 100")
	}

	*fuel = *fuel * (1 + percent/100)

	return nil
}

// FormatTripReport формирует отчет о поездке.
func FormatTripReport(route string, fuel, cost float64) (string, error) {
	if route == "" {
		return "", fmt.Errorf("маршрут не может быть пустым")
	}

	if fuel < 0 || cost < 0 {
		return "", fmt.Errorf("топливо и стоимость не могут быть отрицательными")
	}

	return fmt.Sprintf(
		"Маршрут: %s | Топливо: %.2f л | Стоимость: %.2f тг",
		route,
		fuel,
		cost,
	), nil
}
