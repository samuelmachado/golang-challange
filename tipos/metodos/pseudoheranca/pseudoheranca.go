package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	tudoLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.tudoLigado = true

	fmt.Printf("%s %d %v", f.nome, f.velocidadeAtual, f.tudoLigado)
}
