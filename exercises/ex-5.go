// 5- Escreva uma função que recebe um número inteiro positivo como entrada e calcula o fatorial desse número. O fatorial de um número n é o produto de todos os números inteiros de 1 a n.

package main

import "fmt"

func main(){
	var x int

	fmt.Print("Digite um numero:  ")
	fmt.Scan(&x)

	if x < 0 {
	} else {
		fatorial := 1
		for i := 1; i <= x; i++ {
			fatorial *= i
		}
		fmt.Printf("Fatorial de %d é: %d\n", x, fatorial)
	}
}

