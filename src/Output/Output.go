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