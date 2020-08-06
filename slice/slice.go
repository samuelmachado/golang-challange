package main
import (
	"fmt"
	"reflect"
)
func main() {
	a1 := [3] int{1,2,3} //arr
	s1 := [] int{1,2,3} //slice

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1,2,3,4,5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	
}