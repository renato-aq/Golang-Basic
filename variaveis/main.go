package main

import "fmt"

var (
    nome string = "fulano"
    idade int
    casado bool
)

func main() {
   idade = 32
   casado = true

   carteira := 15.35
   estado_civil := "solteiro"

   if(casado == true) {
       estado_civil = "casado"
   }

   fmt.Println(nome, "tem", idade, "anos de idade e é", estado_civil,".")
   fmt.Println(nome, "tem", carteira, "na carteira.")
}
