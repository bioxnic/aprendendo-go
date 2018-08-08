package main

import "fmt"

func main() {
	// tipos não precisam ser inferidos
	// Não precisa usar var, atribuidor rápido :=
	// Quando não inicializadas o valor atribuido é 0

	nome := "allan"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão do programa: ", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	// fmt.Scanf(&comando) // para inferir o tipo
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Não conheço este comando:", comando)
	// }

	// Não precisa de breaks
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando:", comando)
	}

}
