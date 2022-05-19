package main

import "fmt"

// EXERCICIO 1
func letter(word string) {

	var wordsBag = make(map[string]int)

	for i := 0; i < len(word); i++{
		wordsBag[string(word[i])] += 1
	}
	
	fmt.Printf("A quantidade de letras da palavra %s é %d\n", word, len(word))

	fmt.Printf("As letras contidas em %s são: ", word)
	for index, _ := range wordsBag {
		fmt.Printf("%s ", index)
	}
	fmt.Println()
	fmt.Println()
}

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

func rent(age int, isEmployee bool, experienceYears float64, wage int) {

	var fees bool = false

	if !verifyAge(age) || !verifyEmployment(isEmployee) || !verifyExperience(experienceYears) {
		return
	}

	fees = verifyWage(wage)

	verifyFees(fees)
}

// EXERCÍCIO 3
func monthConverter(monthNumber int){
	var months = map[int]string{
		1: "Janeiro",
		2: "Fevereiro",
		3: "Março",
		4: "Abril",
		5: "Maio",
		6: "Junho",
		7: "Julho",
		8: "Agosto",
		9: "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	if (monthNumber < 13 && monthNumber > 0) {
		fmt.Printf("%d de %s\n", monthNumber, months[monthNumber])
	} else {
		fmt.Println("Mês com esse número não existe")
	}
}

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

// EXERCÍCIO 4
func findEmployee(name string) {
	fmt.Printf("%s is %d years old\n", name, employees[name])
}

func findAge(age int) {
	amount := 0
	for _, age := range employees {
		if age == 21{
			amount++
		}
	}	
	fmt.Printf("A quantidade de funcionários com %d anos é %d.\n", age, amount)
}

func main() {

	letter("Tales")
	fmt.Println()

	rent(23, true, 5, 150000)
	fmt.Println()

	monthConverter(13)
	fmt.Println()

	findEmployee("Benjamin")
	findAge(21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
	fmt.Println()
}