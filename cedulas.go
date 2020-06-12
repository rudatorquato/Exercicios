package main

import (
	"fmt"
)

func main() {
	p := fmt.Println
	var cedulas, cem, result1, cinquenta, result2, cinco int64

	p("Valor:")
	fmt.Scanf("%d", &cedulas)

	if cedulas > 0 {
		cem = cedulas / 100
		p("Notas de 100:")
		p(cem)
	}
	if cem > 0 {
		result1 = cedulas - (cem * 100)
		//p("result1")
		//p(result1)
		if result1 > 0 {
			cinquenta = result1 / 50
			p("Notas de 50:")
			p(cinquenta)
		}
	}
	if cinquenta > 0 {
		result2 = cedulas - (cem * 100) - (cinquenta * 50)
		//p("result2")
		//p(result2)
		if result1 > 0 {
			cinco = result2 / 5
			p("Notas de 5:")
			p(cinco)
		}
	}
}
