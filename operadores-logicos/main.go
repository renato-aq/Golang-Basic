package main

import "fmt"

func main() {
    var (
        valA int8 = 10
        valB int8 = 32
        Op string = ""
    )

    if (valA > valB) {
        fmt.Println("Soma: ", valA + valB)
    } else {
        fmt.Println("subtracao: ", valB - valA)
    }


    switch Op {
    case "+":
        fmt.Println("(switch) soma: ", valA + valB)
    case "-":
        fmt.Println("(switch) subtracao: ", valA - valB)
     default:
         fmt.Println("(switch) multiplicacao: ",int16(valA) * int16(valB))
    }

    fmt.Println("(defer) soma:", valA + valB)
    defer fmt.Println("(defer) subtracao", valA - valB)
     fmt.Println("(defer) multiplicacao", int16(valA) * int16(valB))
    fmt.Println("(defer) divisao", valB % valA)
    
}
