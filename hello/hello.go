package main

import "fmt"
import "reflect"

func main() {
	var nome = "João"
	var idade = 10
	var versao float32 = 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é",idade,"anos.")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}
