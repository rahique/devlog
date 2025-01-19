package main

import (
	"fmt"
	"strings"
)

func conversion(te float64, fr string, to string) float64 {
	var converted float64
	fr = strings.ToUpper(fr)
	to = strings.ToUpper(to)
	if fr == "C" && to == "K" {
		converted = ctok(te)
	} else if fr == "C" && to == "F" {
		converted = ctof(te)
	}
	return converted
}

func ctok(x float64) float64 {
	x += 273
	return x
}
func ctof(x float64) float64 {
	x = (x * (9.0 / 5.0)) + 32
	return x
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

	fmt.Println("yout converted Temp is : ", converted)
}
