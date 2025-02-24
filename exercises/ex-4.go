// 4- Escreva um programa em Go que utilize uma estrutura de dados map para armazenar nomes de pessoas como chaves e suas idades como valores. Imprima o mapa completo.
package main

import (
    "fmt"
)

func main() {
	ni := map[string]int {
		"Bob": 12,
		"Maria": 29,
		"Luis": 40,
	}
	for nome, idade := range ni {
		fmt.Printf("%s: %d\n", nome, idade)
	}

}