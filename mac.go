package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	var mac, MAC string
	var x int

	p("Endereço MAC:")
	fmt.Scanf("%s", &mac)

	MAC = (strings.ToUpper(mac))
	x =strings.Count(mac, "-")

	if MAC == mac {

		if x == 5 {
			p("Endereço MAC verdadeiro")
		}

	}
	
}
//10-C3-7B-C1-F3-D2