package month

import "fmt"

// EXERCÍCIO 3
func MonthConverter(monthNumber int){
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
