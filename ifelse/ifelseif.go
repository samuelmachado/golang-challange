package main
import (
	"fmt"
)

func main() {

	fmt.Println(notaParaConceito(7.3))
	// notaParaConceito(5.1)
}

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10{
		return "A"
	} else if n >= 5 && n < 8 {
		return "B"
	} 
	return "E"
}