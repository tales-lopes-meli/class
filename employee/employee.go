package employee

import "fmt"

// EXERCÍCIO 4

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func FindEmployee(name string) {
	fmt.Printf("%s is %d years old\n", name, employees[name])
}

func FindAge(age int) {
	amount := 0
	for _, age := range employees {
		if age == 21{
			amount++
		}
	}	
	fmt.Printf("A quantidade de funcionários com %d anos é %d.\n", age, amount)
}

func AddEmployee(name string, age int) {
	employees[name] = age
}

func RemoveEmployee(name string) {
	delete(employees, name)
}