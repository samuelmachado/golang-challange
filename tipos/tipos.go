
package main
import (
	"fmt"
	"reflect"
)
func main() {
	fmt.Println(1, 2 , 1000)
	fmt.Println("Literal inteiro", reflect.TypeOf(32000))


	bo:= true
	fmt.Println("Literal inteiro", reflect.TypeOf(bo))

	s1 := "Olá meu nome é Samuel"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `
	Olá
	meu
	nome
	é
	Leo
	`
	fmt.Println("O tamanho da string é", len(s2))
}