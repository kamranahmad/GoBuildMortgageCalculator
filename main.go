package main

import (
	"fmt"
	"math"
)

func main() {
	var principal float64

	fmt.Print("Enter Principal value : ")
	fmt.Scanf("%f", &principal)
	fmt.Printf("Principal You have entered : %f \n", principal)

	var interest float64
	fmt.Print("Enter Yearly Interest Rate: ")
	fmt.Scanf("%f", &interest)
	fmt.Printf("Yearly Interest Rate You have entered : %f \n", interest)

	//monthly interest
	interest = interest / 100 / 12

	var term float64
	fmt.Print("Enter Term (years): ")
	fmt.Scanf("%f", &term)
	fmt.Printf("Years You have entered : %f \n", term)

	//number of months
	term = term * 12

	var payments float64

	payments = ((principal * interest) / (1 - math.Pow(1+interest, -term)))

	fmt.Printf("Your monthly payment is : %f \n", payments)

}
