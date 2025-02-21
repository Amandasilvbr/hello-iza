// 2- Modifique o programa anterior para imprimir um triângulo de asteriscos invertido;
package main

import "fmt"

func main(){
    var l int
   
    fmt.Print("digite a quantidade de linhas:  ")
    fmt.Scan(&l)

    for i := l; i > 0; i-- {
        for j :=0 ; j < i; j++ {
            fmt.Print("*")

        }
        fmt.Println()
    }
   
}