package main

import (
	"fmt"
)

type celsius struct {
	suhu float64
}

type fahrenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

func (c celsius) toCelsius() float64 {
	return c.suhu
}

func (c celsius) toFahrenheit() float64 {
	return (c.suhu * 9.0 / 5.0) + 32
}

func (c celsius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f fahrenheit) toCelsius() float64 {
	return (f.suhu - 32) * 5.0 / 9.0
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

func (f fahrenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * 5.0 / 9.0
}

func (k kelvin) toCelsius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.suhu * 9.0 / 5.0) - 459.67
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

type TemperatureConverter interface {
	toCelsius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func getInput(prompt string) float64 {
	var value float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&value)
		if err == nil {
			return value
		}
		fmt.Println("Input tidak valid. Harap masukkan sebuah angka.")
	}
}

func getMenuChoice(prompt string) int {
	for {
		choice := getInput(prompt)

		intValue := int(choice)

		if float64(intValue) == choice && intValue >= 1 && intValue <= 3 {
			return intValue
		}

		fmt.Println("Pilihan tidak valid. Masukkan angka 1, 2, atau 3.")
	}
}

func main() {
	fmt.Println("--- Program Konversi Suhu ---")

	fmt.Println("\nPilih satuan suhu awal:")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	suhuAwal := getMenuChoice("Masukkan pilihan Anda (1-3): ")

	fmt.Println("\nPilih satuan suhu tujuan:")
	fmt.Println("1. Celsius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	suhuAkhir := getMenuChoice("Masukkan pilihan Anda (1-3): ")

	suhu := getInput("\nMasukkan nilai suhu yang akan dikonversi: ")

	var converter TemperatureConverter
	switch suhuAwal {
	case 1:
		converter = celsius{suhu}
	case 2:
		converter = fahrenheit{suhu}
	case 3:
		converter = kelvin{suhu}
	}

	var suhuFinal float64
	var unit string

	switch suhuAkhir {
	case 1:
		suhuFinal = converter.toCelsius()
		unit = "°C"
	case 2:
		suhuFinal = converter.toFahrenheit()
		unit = "°F"
	case 3:
		suhuFinal = converter.toKelvin()
		unit = "K"
	}

	fmt.Printf("\nHasil konversi: %.2f %s\n", suhuFinal, unit)
}
