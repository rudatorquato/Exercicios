package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	var ipV string

	p("Endereço IP:")
	fmt.Scanf("%s", &ipV)
	v := len(ipV) // se isso for 10.0.0.1  a quantidade de caracteres desse ip equivale ao x
	x := 8 
	//fmt.Println(strings.HasPrefix(ipV, "127"))
	//fmt.Println(strings.HasSuffix(ipV, "255"))

	if v == x {
		p("Novo IP: 192.0.0.1")
	}
	
}

