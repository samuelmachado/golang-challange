package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "notebook", 100.0, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":1, "nome":"Caneca", "preco":9.9, "tags":["Papelaria","Teste"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
