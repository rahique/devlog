package main

import (
	"fmt"
	"strings"
)

// converison function
func conversion(te float64, fr string, to string) float64 {
	var converted float64
	fr = strings.ToUpper(fr)
	to = strings.ToUpper(to)

	if fr == "C" && to == "K" {
		converted = ctok(te)
	} else if fr == "C" && to == "F" {
		converted = ctof(te)
	} else if fr == "K" && to == "C" {
		converted = ktoc(te)
	} else if fr == "K" && to == "F" {
		converted = ktof(te)
	} else if fr == "F" && to == "C" {
		converted = ftoc(te)
	} else if fr == "F" && to == "K" {
		converted = ftok(te)
	} else {
		fmt.Println("Invalid conversion units!")
	}
	return converted
}

// Conversion functions
func ctok(x float64) float64 {
	return x + 273.15
}

func ctof(x float64) float64 {
	return (x * 9.0 / 5.0) + 32
}

func ktoc(x float64) float64 {
	return x - 273.15
}

func ktof(x float64) float64 {
	return (x-273.15)*9.0/5.0 + 32
}

func ftoc(x float64) float64 {
	return (x - 32) * 5.0 / 9.0
}

func ftok(x float64) float64 {
	return (x-32)*5.0/9.0 + 273.15
}

func main() {
	var temp float64 = 0
	var funit, tounit string
	fmt.Print("Enter Temp : ")
	fmt.Scan(&temp)
	fmt.Print("Enter Unit : ")
	fmt.Scan(&funit)
	fmt.Print("Enter Unit to Convert : ")
	fmt.Scan(&tounit)

	var converted float64 = conversion(temp, funit, tounit)

	fmt.Printf("yout converted Temp is : %.2f \n", converted)
}
