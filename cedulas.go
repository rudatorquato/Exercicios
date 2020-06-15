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
	}
	if cem > 0 {
		result1 = cedulas - (cem * 100)
		if result1 > 0 {
			cinquenta = result1 / 50
		}
	}
	if cinquenta > 0 {
		result2 = cedulas - (cem * 100) - (cinquenta * 50)
		if result1 > 0 {
			cinco = result2 / 5
		}		
	}
	p("Notas de 100:", cem)
		p("Notas de 50:", cinquenta)
		p("Notas de 5:",cinco)
}
