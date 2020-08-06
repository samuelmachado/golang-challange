package main
import (
	"fmt"
)


func main(){
	fmt.Println(tipo(2))
	fmt.Println(tipo("OPA"))

}
func tipo(i interface{}) string {
	switch i.(type) {
		case int: return "inteiro"
		case func(): return "função"
		default: return "não sei"
	}
}