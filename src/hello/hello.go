package main

import "fmt"
import "reflect"

func main() {
	// tipos não precisam ser inferidos
	// Não precisa usar var, atribuidor rápido :=
	// Quando não inicializadas o valor atribuido é 0

	nome := "allan"
	idade := 18
	versao := 1.1
	fmt.Println("Olá", nome, ", sua idade é", idade)
	fmt.Println("Versão: ", versao)
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}
