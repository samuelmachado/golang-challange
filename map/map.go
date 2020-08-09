package main
import (
	"fmt"
	// "reflect"
)
func main() {
	aprovados := make(map[int] string)
	aprovados[12345678] = "Maria"
	aprovados[12345670] = "José"

	for cpf, nome := range aprovados {
		fmt.Println(nome, cpf)
	}

	fmt.Println(aprovados)
	delete(aprovados, 12345670)
	for cpf, nome := range aprovados {
		fmt.Println(nome, cpf)
	}
}
