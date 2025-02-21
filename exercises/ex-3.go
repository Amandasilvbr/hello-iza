// 3- Crie uma função personalizada que recebe um número inteiro positivo como entrada. A função deve retornar o resultado de uma operação qualquer realizada com esse número. No entanto, se o número de entrada for negativo, a função deve retornar um erro indicando que um número negativo foi fornecido.
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scan(&n)

    if n < 0 {
        fmt.Println("Número negativo")
    } else {
        resultado := n * n 
        fmt.Printf("O resultado da operação (n²) é: %d\n", resultado)
    }
}