package main

import (
	m "math"
	"fmt"
)

func main() {
	const PI float64 = 3.1415
	raio := 3.2
	area := PI * m.Pow(raio,2)
	fmt.Println("Área da circunferência é", area)

	const(
		a = 1
		b = 2
	)

	var(
		c = 1
		e = 2
	)

	fmt.Println(a,b,c,e)

	
}