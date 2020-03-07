package agasparo

import (
	"github.com/fatih/color"
	"fmt"
	"maths_42"
)

type Equation struct {

	Puis0 float64
	Puis1 float64
	Puis2 float64
}

func Array_search(array []string, to_search string) (res int) {

	for i:= 0; i < len(array); i++ {

		if array[i] == to_search {
			return (i)
		}
	}
	return (-1)
}

func Array_search_count(array []string, to_search string) (res int) {

	count := 0

	for i:= 0; i < len(array); i++ {

		if array[i] == to_search || array[i][0] == to_search[0] {
			count++;
		}
	}
	return (count)
}

func Delta(Eq Equation) {

	fmt.Println("Etape 1 :")
	color.Green("Δ = b² - 4ac")
	fmt.Printf("a = %f\n", Eq.Puis2)
	fmt.Printf("b = %f\n", Eq.Puis1)
	fmt.Printf("c = %f\n", Eq.Puis0)
	fmt.Printf("Δ = %f² - 4 * %f * %f\n", Eq.Puis1, Eq.Puis2, Eq.Puis0)
	b := maths_42.Pow(Eq.Puis1, 2)
	other := 4 * Eq.Puis2 * Eq.Puis0
	delta := b - other
	fmt.Printf("Δ = %f - %f\n", b, other)
	color.Green("Δ = %f", delta)

	if delta < 0 {
		color.Green("Δ < 0, No solutions for this equation")
		return
	}

	if delta == 0 {
		color.Green("Δ == 0, One solution : x0 = -b / 2a")
		DeltaNil(Eq, delta)
		return
	}

	if delta > 0 {
		color.Green("Δ > 0, 2 solutions : x1 = (- b + √Δ) / 2a and x2 = (- b - √Δ) / 2a")
		DeltaSup(Eq, delta)
		return
	}
}

func DeltaNil(Eq Equation, delta float64) {
	
	fmt.Println("Etape 2 :")
	fmt.Printf("x0 = - %f / 2 * %f\n", Eq.Puis1, Eq.Puis2)
	b := Eq.Puis1 * -1
	other := Eq.Puis2 * 2
	fmt.Printf("x0 = %f / %f\n", b, other)
	res := b / other
	fmt.Printf("x0 = %f\n", res)
	Rational := maths_42.Rational{res, 0, 0, "", 3, ""}
	maths_42.Trasnform(&Rational)
	color.Set(color.FgGreen)
		fmt.Printf("Resulat :\n 	 x0 = %f ou x0 = ", res)
		fmt.Println(Rational.Frac)
	color.Unset()
}

func DeltaSup(Eq Equation, delta float64) {

	fmt.Println("Etape 2 :")
	fmt.Printf("x1 = (- %f + √%f) / 2 * %f and x2 = (- %f - √%f) / 2 * %f\n", Eq.Puis1, delta, Eq.Puis2, Eq.Puis1, delta, Eq.Puis2)
	deb := (Eq.Puis1 * -1)
	fin := (Eq.Puis2 * 2)
	fmt.Printf("x1 = (%f + √%f) / %f and x2 = (%f - √%f) / %f\n", deb, delta, fin, deb, delta, fin)
	deb_x1 := deb + Sqrt(delta)
	deb_x2 := deb - Sqrt(delta)
	fmt.Printf("x1 = %f / %f and x2 = %f / %f\n", deb_x1, fin, deb_x2, fin)
	fin_x1 := deb_x1 / fin
	fin_x2 := deb_x2 / fin
	fmt.Printf("x1 = %f and x2 = %f\n", fin_x1, fin_x2)
	color.Set(color.FgGreen)
	if isFloatInt(Sqrt(delta)) {
		fmt.Printf("Resultat : \n 	x1 = %f ou x1 = ", fin_x1)
		Rationalx1 := maths_42.Rational{fin_x1, 0, 0, "", 3, ""}
		maths_42.Trasnform(&Rationalx1)
		fmt.Println(Rationalx1.Frac)
		fmt.Printf("	x2 = %f ou x2 = ", fin_x2)
		Rationalx2 := maths_42.Rational{fin_x2, 0, 0, "", 3, ""}
		maths_42.Trasnform(&Rationalx2)
		fmt.Println(Rationalx2.Frac)
	} else {
		fmt.Printf("Resultat : \n 	x1 = %f ou ", fin_x1)
		Rationalx1 := fmt.Sprintf("x1 = (%f + √%f) / %f", deb, delta, fin)
		fmt.Println(Rationalx1)
		fmt.Printf("	x2 = %f ou ", fin_x2)
		Rationalx2 := fmt.Sprintf("x2 = (%f - √%f) / %f", deb, delta, fin)
		fmt.Println(Rationalx2)
	}
	color.Unset()
}

func Deg1(Eq Equation) {

	termD := Inverse(Eq.Puis0)
	termG := Eq.Puis1

	fmt.Println("Etape 1 :")
	fmt.Printf("%f * X = %f\n", termG, termD)
	fmt.Println("Etape 2 :")
	fmt.Printf("X = %f / %f\n", termD, termG)
	res := termD / termG
	fmt.Println("Etape 3 :")
	fmt.Printf("X = %f\n", res)
	Rational := maths_42.Rational{res, 0, 0, "", 3, ""}
	maths_42.Trasnform(&Rational)
	color.Set(color.FgGreen)
		fmt.Printf("Resulat :\n 	 X = %f ou X = ", res)
		fmt.Println(Rational.Frac)
	color.Unset()
}

func Inverse(nb float64) (res float64) {

	return (nb * -1)
}

func isFloatInt(floatValue float64) bool {

    return floatValue == float64(int(floatValue))
}

func Sqrt(x float64) float64 {
    
    var z float64 = 1

    for i := 1; i <= 10; i++ {

    	z = (z - (maths_42.Pow(z, 2) - x) / ( 2 * z))
    }
    return (z)
}