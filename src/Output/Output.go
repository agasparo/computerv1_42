package Output

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func PrintErrorSyntax(str string, err string) {

	explode := strings.Split(str, err)
	fmt.Printf("	%s", explode[0])
	color.Set(color.FgRed)
	fmt.Printf(err)
	color.Unset()
	fmt.Println(explode[1])
}

func PrintNewEquation(p2 float64, p1 float64, p0 float64) {

	color.Set(color.FgGreen)
	fmt.Print("New equation : ")
	if (p2 != 0) {
		fmt.Printf("%f * X^2", p2)
		fmt.Print(" + ")
	}
	if (p1 != 0) {
		fmt.Printf("%f * X^1", p1)
		fmt.Print(" + ")
	}
	if (p0 != 0) {
		fmt.Printf("%f * X^0", p0)
	}
	fmt.Println(" = 0")
	color.Unset()
}

func Usage() {
	fmt.Println("Computerv1 : Usage : [-natural=(1/0)] [equation]")
	fmt.Println("Equation must be like : a * X^nb + b * X^nb + c * X^nb = d * X^nb + e * X^nb + f * X^nb")
	fmt.Println("The power (nb) must be an int between 0 and 2")
	fmt.Println("Exemple : '5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0'")
}