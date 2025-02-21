// 1- Escreva um programa em Go que imprima um triângulo de asteriscos de acordo com o número de linhas que o usuário desejar;
package main

import "fmt"

func main(){
    var l int
   
    fmt.Print("digite a quantidade de linhas:  ")
    fmt.Scan(&l)

    for i := 1; i <=l; i++ {
       
        for j := 0; j < i; j++ {
            fmt.Print("*")

        }
        fmt.Println()
    }
   
}
