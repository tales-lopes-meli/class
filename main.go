package main

import (
	"github.com/tales-lopes-meli/class/letter"
	"github.com/tales-lopes-meli/class/employee"
	"github.com/tales-lopes-meli/class/rent"
	"github.com/tales-lopes-meli/class/month"
)

func main() {

	var word string = "Carambola"
	letter.CountLetters(word)

	var (
		age = 23
		isEmployee = true
		experienceYears = 1.5
		wage = 150000
	)
	rent.VerifyRent(age, isEmployee, experienceYears, wage)

	var monthNumber int = 11
	month.MonthConverter(monthNumber)

	var name string = "Benjamin"
	employee.FindEmployee(name)
	
	age = 21
	employee.FindAge(age)

	name = "Federico"
	age = 25
	employee.AddEmployee(name, age)
	
	name = "Pedro"
	employee.RemoveEmployee(name)
}