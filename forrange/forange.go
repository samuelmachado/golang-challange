package main
import (
	"fmt"
)
func main() {
	numeros := [...] int{1,2,3,4,5}

	for i, numero := range numeros {
		fmt.Println("Res", i, numero)
	}

	for _, numero := range numeros {
		fmt.Println("Res", numero)
	}
}