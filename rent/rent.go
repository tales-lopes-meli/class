package rent

import "fmt"

// EXERCICIO 2

func verifyAge(age int) bool {
	if age <= 22 {
		fmt.Println("Desculpa! Empréstimos somente são autorizados a maiores de 22 anos")
		return false
	}
	return true
}

func verifyEmployment(isEmployee bool) bool {
	if !isEmployee {
		fmt.Println("Desculpa! Empréstimos somente são autorizados a pessoas empregadas")
		return false
	}
	return true
}

func verifyExperience(experienceYears float64) bool {
	if experienceYears <= 1{
		fmt.Println("Desculpa! Empréstimos somente são autorizados a pessoas com mais de um ano de experiência")
		return false
	}
	return true
}

func verifyWage(wage int) bool {
	if wage <= 100000 {
		return true
	}
	return false
}

func verifyFees(fees bool) {
	if fees {
		fmt.Println("Parabéns! Seu empréstimo foi aprovado com juros")
	} else {
		fmt.Println("Parabéns! Seu empréstimo foi aprovado sem juros")
	}
}

func VerifyRent(age int, isEmployee bool, experienceYears float64, wage int) {

	var fees bool = false

	if !verifyAge(age) || !verifyEmployment(isEmployee) || !verifyExperience(experienceYears) {
		return
	}

	fees = verifyWage(wage)

	verifyFees(fees)
}