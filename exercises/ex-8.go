// 8- Escreva uma função que recebe uma lista de números inteiros como entrada e retorna o valor máximo e mínimo encontrados nessa lista.

package main

import (
	"fmt"
)

func main() {
	var numeros [5]int 

	fmt.Println("Digite 5 números inteiros:  ")

	for i := 0; i < 5; i++ {
		fmt.Printf("Número %d: ", i+1) 
		fmt.Scan(&numeros[i]) 
	}

	min := numeros[0]
	max := numeros[0]

	for _, n := range numeros {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Printf("Máximo: %d, Mínimo: %d\n", max, min)
}