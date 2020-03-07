package maths_42

import (
	"fmt"
)

type Rational struct {

	Nb	 	float64
	Num 	int64
	Den 	int64
	Sign	string
	Preci	int
	Frac	string
}

func Trasnform(Ra *Rational) {
	
	var mul = 0
	var new_nb int64

	if isuseless(Ra.Nb) {
		return
	}

	if Ra.Nb > 0 {
		Ra.Sign = "+"
	} else {
		Ra.Sign = "-"
		Ra.Nb *= -1 
	}

	if int(Ra.Nb) > 0 {
		mul = getMult(Ra.Nb - float64(int(Ra.Nb)))
	} else {
		mul = getMult(Ra.Nb)
	}

	deno := int64(Pow(10, mul + Ra.Preci))
	new_nb = int64((Ra.Nb * float64(deno)))
	Processus(new_nb, deno, Ra)
}

func Processus(new_nb, deno int64, Ra *Rational) {
	
	for pgcd := gcf(new_nb, deno); pgcd != 1; pgcd = gcf(new_nb, deno) {

		new_nb /= pgcd
		deno /= pgcd
	}
	Ra.Num = new_nb
	Ra.Den = deno
	if Ra.Sign == "-" {
		Ra.Frac += Ra.Sign
	}
	Ra.Frac += fmt.Sprintf("%d/%d", Ra.Num, Ra.Den) 
}

func gcf(a, b int64) int64 {

	if a < b {
        return gcf(b, a)
    }

    if b == 0 {
        return a
    }

    a = a % b
    return gcf(b, a)
}

func getMult(nb float64) (count int){

	for new_nb := nb; int(new_nb) <= 0; new_nb *= 10 {
		count++
	}

	return (count)
}

func Pow(x float64, n int) (res float64) {

    number := 1.00;

    for i := 0; i < n; i++ {
        number *= x;
    }

    return (number);
}

func isuseless(floatValue float64) bool {

    return floatValue == float64(int(floatValue))
}