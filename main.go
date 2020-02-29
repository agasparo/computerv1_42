package main

import (
	"fmt"
	"check"
	"github.com/fatih/color"
	"strings"
	"Output"
    "agasparo"
)

func main() {

    // check le double =
    // check qu'il ya toujouts x^0, 1 ou 2

    c := "42 * X^0 = 42 * X^0"

    fmt.Println("[1/4] Init Check struct ...")
    Check_struct := Init(c)
    fmt.Println("done")

    index := agasparo.Array_search(Check_struct.Tab_str, "=")
    if index < 0 {
        color.Red("Notice: %s", "you must have an = on your equation")
        return
    }

    fmt.Println("[2/4] Init Struct eq 1 ...")
    Eq1 := check.CheckPuis{}
    if check.SetPuis(&Check_struct, &Eq1, 0) == 0 {
    	color.Red("Notice: %s on caractere %d ", Check_struct.Error, Check_struct.Error_value)
    	Output.PrintErrorSyntax(c, Check_struct.Error_desc)
    	return 
    }

    fmt.Println("done")
    fmt.Println("[3/4] Init Struct eq 2 ...")
    Eq2 := check.CheckPuis{}
    if check.SetPuis(&Check_struct, &Eq2, index + 1) == 0 {
    	color.Red("Notice: %s on caractere %d ", Check_struct.Error, Check_struct.Error_value)
    	Output.PrintErrorSyntax(c, Check_struct.Error_desc)
    	return 
    }

    fmt.Println("done")
    fmt.Println("[4/4] Calculating final Struct ...")
    EqF := agasparo.Equation{
        Eq1.Puis0 + agasparo.Inverse(Eq2.Puis0),
        Eq1.Puis1 + agasparo.Inverse(Eq2.Puis1),
        Eq1.Puis2 + agasparo.Inverse(Eq2.Puis2),
    }

    if EqF.Puis2 != 0 {
        Output.PrintNewEquation(EqF.Puis2, EqF.Puis1, EqF.Puis0)
        fmt.Println("Equation degree : 2")
        agasparo.Delta(EqF)
    } else if EqF.Puis1 != 0 {
        Output.PrintNewEquation(EqF.Puis2, EqF.Puis1, EqF.Puis0)
        fmt.Println("Equation degree : 1")
        agasparo.Deg1(EqF)
    } else {
        fmt.Println("Equation degree : 0")
        color.Green("All real numbers are solution")
    }
}

func Init(str string) (check.CheckParam) {

	init_struct := check.CheckParam{}
	new_str := strings.Trim(str, " ")
	init_struct.Tab_str = strings.Split(new_str, " ")
	return (init_struct)
}