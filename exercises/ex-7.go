// 7- Escreva uma função que recebe uma string como entrada e retorna uma nova string que é a versão invertida da original.

package main

import "fmt"

func main() {
    var a string

    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&a)

    invertida := "" 

    for i := len(a) - 1; i >= 0; i-- {
        invertida += string(a[i]) 
    }
	
    fmt.Println("Palavra invertida:", invertida)
}