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
func rent(age int, isEmployee bool, experienceYears float64, wage int) {

	var fees bool = false
	var approved bool = true

	if age <= 22 {
		fmt.Println("Desculpa! Empréstimos somente são autorizados a maiores de 22 anos")
		approved = false
	}
	if !isEmployee {
		fmt.Println("Desculpa! Empréstimos somente são autorizados a pessoas empregadas")
		approved = false
	}
	if experienceYears <= 1{
		fmt.Println("Desculpa! Empréstimos somente são autorizados a pessoas com mais de um ano de experiência")
		approved = false
	}

	if !approved {
		return
	}

	if wage <= 100000 {
		fees = true
	}

	if fees {
		fmt.Println("Parabéns! Seu empréstimo foi aprovado com juros")
	} else {
		fmt.Println("Parabéns! Seu empréstimo foi aprovado sem juros")
	}
	fmt.Println()
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

// EXERCÍCIO 4
func findEmployee(name string) {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("%s is %d years old\n", name, employees[name])
}

func main() {

	letter("Tales")

	rent(23, true, 5, 100000)

	monthConverter(13)

	findEmployee("Benjamin")
}