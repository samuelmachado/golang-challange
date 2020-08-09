package main

import "fmt"

type esporitvo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esporitvo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo....")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")

}
func main() {
	var b esportivoLuxuoso = bmw7{}
	b.fazerBaliza()
	b.ligarTurbo()
}
