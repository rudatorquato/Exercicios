package main
import (
	"fmt"
)

func main() {
	//var x, y int
	var x int = 2
	var y int = 10
	p := fmt.Println

	p("Informe o numero:\n")
	//var primo bool

	for i := x; i < y+1; i++ {
		//primo = true

		for j := 2; j < i; j++ {

			if i%j == 0 {
				//primo = false
				break
			}
		}
		

		//if primo {
			p(i)
		//}
		if i == 2 {
			p("Numero 2: Frase")
		}else if i == 3 {
			p("Numero 3: Frase")
		}else if i == 5 {
			p("Numero 5: Frase")
		}else if i == 7 {
			p("Numero 7: Frase")
		}else if i == 11 {
			p("Numero 11: Frase")
		}else if i == 13 {
			p("Numero 33: Frase")
		}else if i == 17 {
			p("Numero 17: Frase")
		}else if i == 19 {
			p("Numero 19: Frase")
		}else if i == 23 {
			p("Numero 23: Frase")
		}else if i == 29 {
			p("Numero 29: Frase")
		}else if i == 31 {
			p("Numero 31: Frase")
		}else if i == 37 {
			p("Numero 37: Frase")
		}else if i == 41 {
			p("Numero 41: Frase")
		}else if i == 43 {
			p("Numero 43: Frase")
		}else if i == 47 {
			p("Numero 47: Frase")
		}else if i == 53 {
			p("Numero 53: Frase")
		}else if i == 59 {
			p("Numero 59: Frase")
		}else if i == 61 {
			p("Numero 61: Frase")
		}else if i == 67 {
			p("Numero 67: Frase")
		}else if i == 71 {
			p("Numero 71: Frase")
		}else if i == 73 {
			p("Numero 73: Frase")
		}else if i == 79 {
			p("Numero 79: Frase")
		}else if i == 83 {
			p("Numero 83: Frase")
		}else if i == 89 {
			p("Numero 89: Frase")
		}else if i == 97 {
			p("Numero 97: Frase")
		}
		
		
	}
}