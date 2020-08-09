package main
import (
	"fmt"
)

func media(numeros ... float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println("Média", media(1.1, 22.2, 5.0,9.9))
}