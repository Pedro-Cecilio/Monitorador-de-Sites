package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	qtMonitoramento = 2
	tempoEspera     = 10
)

func main() {
	exibeIntruducao()
	for {

		exibeMenu()

		comando := leComando()

		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("Exibindo Logs...")
			imprimeLogs()

		} else if comando == 0 {
			fmt.Println("Saindo...")
			os.Exit(0)
		} else {
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}

}

func exibeIntruducao() {
	nome := "Pedro"
	versao := 1.1

	fmt.Println("Olá, sr.", nome+".")
	fmt.Println("Versão do Programa:", versao)
}
func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair.")
}
func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	return comando
}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < qtMonitoramento; i++ {
		for _, url := range sites {
			testaSite(url)
			fmt.Println("")
		}
		time.Sleep(tempoEspera * time.Second)
		fmt.Println("")
	}
}
func testaSite(url string) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode == 200 {
		fmt.Println("O site", url, "foi carregado com sucesso!")
		registraLog(url, true)
	} else {
		fmt.Println("O site", url, "está com problemas. Status code:", res.StatusCode)
		registraLog(url, false)
	}
}
func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool((status)) + "\n")

	arquivo.Close()
}
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
