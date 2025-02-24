// 6- Escreva uma função que recebe uma string como entrada e conta o número de vogais presentes nessa string. Considere vogais apenas as letras "a", "e", "i", "o" e "u".

package main

import "fmt"

func main() {
    var b string
    vogais := 0

    fmt.Print("Escreva uma palavra: ")
    fmt.Scan(&b)

    for _, char := range b {
        switch char {
        case 'a', 'e', 'i', 'o', 'u':
            vogais++
        }
    }

    fmt.Printf("Quantidade de vogais:  %d\n", vogais)
}