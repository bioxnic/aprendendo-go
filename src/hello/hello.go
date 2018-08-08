package main

import "fmt"

func main() {
	// tipos não precisam ser inferidos
	// Não precisa usar var, atribuidor rápido :=
	// Quando não inicializadas o valor atribuido é 0

	nome := "allan"
	// idade := 18
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão do programa: ", versao)

	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O ponteiro/endereço da variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
