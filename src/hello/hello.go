package main

import "fmt"
import "os"
import "net/http"

func main() {

	exibeIntroducao()

	exibeMenu()
	//ignorando primeira var
	_, idade := devolveNomeEIdade()

	fmt.Println("idade", idade)
	comando := leComando()
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
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando:", comando)
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	// tipos não precisam ser inferidos
	// Não precisa usar var, atribuidor rápido :=
	// Quando não inicializadas o valor atribuido é 0

	nome := "allan"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão do programa: ", versao)
}
func leComando() int {
	var comandoLido int
	// fmt.Scanf(&comando) // para inferir o tipo
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	site := "http://www.alura.com.br"
	resp, _ := http.Get(site)

	fmt.Println(resp)
}

func devolveNomeEIdade() (string, int) {
	nome := "Allan"
	idade := 18
	return nome, idade
}
