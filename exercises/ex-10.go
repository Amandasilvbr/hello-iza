// 10- Escreva uma função que recebe um número como entrada e verifica se ele é positivo, negativo ou zero. A função deve retornar uma mensagem indicando o resultado.


package main

import "fmt"

func numero(a int) string {
	switch {
	case a > 0:
		return "Positivo."
	case a < 0:
		return "Negativo."
	default:
		return "Zero."
	}
}

func main() {
	var a int

	fmt.Print("Digite um número: ")
	fmt.Scan(&a)

	resultado := numero(a)
	fmt.Println(resultado)
}