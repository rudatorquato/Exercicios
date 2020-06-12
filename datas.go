package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	p := fmt.Println
	var ano, dia, ano2, dia2 int
	var mes, mes2 string

	p("ANO DATA 1")
	fmt.Scanf("%d", &ano)

	p("mes DATA 1")
	fmt.Scanf("%s", &mes)

	p("dia DATA 1")
	fmt.Scanf("%d", &dia)

	p("ANO DATA 2")
	fmt.Scanf("%d", &ano2)

	p("mes DATA 2")
	fmt.Scanf("%s", &mes2)

	p("dia DATA 2")
	fmt.Scanf("%d", &dia2)

	var m, _ = strconv.Atoi(mes)
	var x = time.Month(m)
	var m2, _ = strconv.Atoi(mes2)
	var x2 = time.Month(m2)
	//p(m)
	//p(x)

	//now := time.Now()
	now := time.Date(
		ano2, x2, dia2, 20, 34, 58, 651387237, time.UTC)
	//p(now)

	then := time.Date(
		ano, x, dia, 20, 34, 58, 651387237, time.UTC)
	//p(then)

	//p(then.After(now))

	if then.After(now) == false {
		p(now)
	} else {
		p(then)
	}

}
