package letter

import "fmt"

// EXERCICIO 1
func CountLetters(word string) {

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