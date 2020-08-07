package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)

	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{
				produtoID: 1,
				qtde:      3,
				preco:     50,
			},
		},
	}
	fmt.Println(pedido.valorTotal())
}
