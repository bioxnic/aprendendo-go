package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	// exibeNomes()
	exibeIntroducao()

	for {

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
	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Site", i, ":", site)
			testaSite(site)
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
	}

	fmt.Println("")
}

func devolveNomeEIdade() (string, int) {
	nome := "Allan"
	idade := 18
	return nome, idade
}

// Quando atingir o limite da capacidade, ela é dobrada
// No fundo o slice é um array
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("Nomes:", nomes)
	fmt.Println("Tamanho do slice:", len(nomes))
	fmt.Println("Capacidade slice:", cap(nomes))

	nomes = append(nomes, "Aparecida")

	fmt.Println("Nomes:", nomes)
	fmt.Println("Tamanho do slice:", len(nomes))
	fmt.Println("Capacidade slice:", cap(nomes))

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code", resp.StatusCode)
	}
}
