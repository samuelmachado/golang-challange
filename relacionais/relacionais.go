package main
import (
	"fmt"
	"time"
)
func main() {

	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)
	fmt.Println("Data:", d1==d2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	fmt.Println("Pessoas", p1)
}
