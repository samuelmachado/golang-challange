package main
import (
	"fmt"
	"strconv"
)
func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))


	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println(strconv.Atoi("123"))

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}