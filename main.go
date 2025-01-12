package main

import (
	"fmt"
)

type celcius struct {
	suhu float64
}

type farenheit struct {
	suhu float64
}

type kelvin struct {
	suhu float64
}

func (c celcius) toCelcius() float64 {
	return c.suhu
}

func (c celcius) toFarenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32
}

func (c celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f farenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

func (f farenheit) toFarenheit() float64 {
	return f.suhu
}

func (f farenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFarenheit() float64 {
	return (k.suhu * (9.0 / 5.0)) - 459.67
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

type hitungSuhu interface {
	toCelcius() float64
	toFarenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Masukkan opsi suhu awal:")
	fmt.Println("1. Celcius")
	fmt.Println("2. Farenheit")
	fmt.Println("3. Kelvin")

	var suhuAwal int
	fmt.Scanf("%d", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 3 {
		fmt.Println("Pilihan tidak valid. Masukkan angka antara 1 dan 3.")
		fmt.Scanf("%d", &suhuAwal)
	}

	fmt.Println("Masukkan opsi suhu akhir:")
	fmt.Println("1. Celcius")
	fmt.Println("2. Farenheit")
	fmt.Println("3. Kelvin")

	var suhuAkhir int
	fmt.Scanf("%d", &suhuAkhir)
	for suhuAkhir < 1 || suhuAkhir > 3 {
		fmt.Println("Pilihan tidak valid. Masukkan angka antara 1 dan 3.")
		fmt.Scanf("%d", &suhuAkhir)
	}

	var suhu float64
	fmt.Println("Masukkan suhu:")
	_, err := fmt.Scanf("%f", &suhu)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	var interfaceSuhu hitungSuhu
	switch suhuAwal {
	case 1:
		interfaceSuhu = celcius{suhu}
	case 2:
		interfaceSuhu = farenheit{suhu}
	case 3:
		interfaceSuhu = kelvin{suhu}
	}

	var suhuFinal float64
	switch suhuAkhir {
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
	case 2:
		suhuFinal = interfaceSuhu.toFarenheit()
	case 3:
		suhuFinal = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Suhu yang dikonversi: %.2f\n", suhuFinal)
}
