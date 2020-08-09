package main
import (
	"fmt"
	// "reflect"
)
func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 12354.1,
			"Guga Pereira": 3.1,
			
		},
		"J": {
			"Gabriela Silva": 12354.1,
			"Guga Pereira": 3.1,	
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range(funcsPorLetra) {
		fmt.Println(letra, funcs)
	}
}

