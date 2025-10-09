/*
 * @author Bianca Bertinato
 * @version 1.0.0
 * @date 2025-10-08
 * @fileoverview This program calculates the gross pay for an employee, who worked 40 hours and had an hourly wage of $8.20.
 */

 package main

 import "fmt"

 func main() { 
	 //shows conversion
	 fmt.Println("What is the gross pay for an employee named Fred, who worked 40 hours at an hourly wage of $8.20?")
	 fmt.Println("Gross pay = " + " $" + fmt.Sprint(40 * 8.20))

	 fmt.Println("\nDone.")
 }