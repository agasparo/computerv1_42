package main

import (
	"os/exec"
	"fmt"
	"log"
	"bytes"
	"strings"
	"github.com/fatih/color"
)

type Testes struct {
	
	Name		string
	FileName	string
	Commande	string
	Outpout		string
	Type_search	string
}

func main() {

	AllTests := make(map[int]Testes)

	AddTest(AllTests, "syntax error 1 (double egale)", "/home/user42/Bureau/computerv1/main.go", "2 * X^0 == 1 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 2 (twice = on equation)", "/home/user42/Bureau/computerv1/main.go", "2 * X^0 = 1 * X^1 = 2 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 3 (L for X)", "/home/user42/Bureau/computerv1/main.go", "2 * X^0 = 1 * L^1", "No solution", "regex")
	AddTest(AllTests, "syntax error 4 (s for ^)", "/home/user42/Bureau/computerv1/main.go", "2 * Xs0 = 1 * X^1", "0", "regex")
	AddTest(AllTests, "syntax error 5 (no space)", "/home/user42/Bureau/computerv1/main.go", "2*X^1.02=1*X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 6 (power not int)", "/home/user42/Bureau/computerv1/main.go", "2 * X^1.02 = 1 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 7 (power alpha)", "/home/user42/Bureau/computerv1/main.go", "2 * X^a = 1 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 8 (terme alpha)", "/home/user42/Bureau/computerv1/main.go", "a * X^1 = 1 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 9 (+ for *)", "/home/user42/Bureau/computerv1/main.go", "2 + X^1 = 1 * X^1", "No solution", "regex")
	AddTest(AllTests, "syntax error 10 (no ^)", "/home/user42/Bureau/computerv1/main.go", "2 * X1 = 1 * X1", "0", "regex")
	AddTest(AllTests, "syntax error 11 (power > 2)", "/home/user42/Bureau/computerv1/main.go", "2 * X^3 = 1 * X^1", "Notice", "regex")
	AddTest(AllTests, "syntax error 12 (power < 0)", "/home/user42/Bureau/computerv1/main.go", "2 * X^1 = 1 * X^-1", "Notice", "regex")

	AddTest(AllTests, "good eq 1 : delta = 0", "/home/user42/Bureau/computerv1/main.go", "2 * X^2 - 3 * X^1 + 9/8 * X^0 = 0", "Δ == 0, One solution", "regex")
	AddTest(AllTests, "good eq 1 : delta = 0", "/home/user42/Bureau/computerv1/main.go", "6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1", "Δ == 0, One solution", "regex")

	AddTest(AllTests, "good eq 2 : delta > 0", "/home/user42/Bureau/computerv1/main.go", "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0", "Δ > 0, 2 solutions", "regex")
	AddTest(AllTests, "good eq 2 : delta > 0", "/home/user42/Bureau/computerv1/main.go", "5 * X^0 + 13 * X^1 + 3 * X^2 = 1 * X^0 + 1 * X^1", "Δ > 0, 2 solutions", "regex")

	AddTest(AllTests, "good eq 3 : delta < 0", "/home/user42/Bureau/computerv1/main.go", "1 * X^2 = -3 * X^1 - 10 * X^0", "Δ < 0, No solutions for this equation", "regex")
	AddTest(AllTests, "good eq 3 : delta < 0", "/home/user42/Bureau/computerv1/main.go", "5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1", "Δ < 0, No solutions for this equation", "regex")

	AddTest(AllTests, "good eq deg1 : 1", "/home/user42/Bureau/computerv1/main.go", "5 * X^0 + 4 * X^1 = 4 * X^0", "-0.25", "regex")
	AddTest(AllTests, "good eq deg1 : 2", "/home/user42/Bureau/computerv1/main.go", "5 * X^0 = 4 * X^0 + 7 * X^1", "0.1428", "regex")

	AddTest(AllTests, "eq special : 1", "/home/user42/Bureau/computerv1/main.go", "42 * X^0 = 42 * X^0", "All real numbers are solution", "regex")
	AddTest(AllTests, "eq special : 2", "/home/user42/Bureau/computerv1/main.go", "42 * X^0 = 4 * X^0", "No solution", "regex")
	
	Header()
	Run(AllTests)
}

func Header() {

	fmt.Println("\n\n")
	color.Magenta(" ::::::::   ::::::::  ::::    ::::  :::::::::  :::    ::: ::::::::::: :::::::::: :::::::::       :::     :::   :::   ")
	color.Magenta(":+:    :+: :+:    :+: +:+:+: :+:+:+ :+:    :+: :+:    :+:     :+:     :+:        :+:    :+:      :+:     :+: :+:+:   ")
	color.Magenta("+:+        +:+    +:+ +:+ +:+:+ +:+ +:+    +:+ +:+    +:+     +:+     +:+        +:+    +:+      +:+     +:+   +:+   ")
	color.Magenta("+#+        +#+    +:+ +#+  +:+  +#+ +#++:++#+  +#+    +:+     +#+     +#++:++#   +#++:++#:       +#+     +:+   +#+   ")
	color.Magenta("+#+        +#+    +#+ +#+       +#+ +#+        +#+    +#+     +#+     +#+        +#+    +#+       +#+   +#+    +#+   ")
	color.Magenta("#+#    #+# #+#    #+# #+#       #+# #+#        #+#    #+#     #+#     #+#        #+#    #+#        #+#+#+#     #+#   ")
	color.Magenta(" ########   ########  ###       ### ###         ########      ###     ########## ###    ###          ###     ####### ")
	color.Magenta("\n\nMain test by : Agasparo\n\n")
}

func AddTest(table map[int]Testes, name string, filename string, Commande string, Outpout string, Type string) {

	NewTest := Testes{
		name,
		filename,
		Commande,
		Outpout,
		Type,
	}
	table[len(table)] = NewTest
}

func Run(table map[int]Testes) {

	for i := 0; i < len(table); i++ {

		cmd := exec.Command("go", "run", table[i].FileName, table[i].Commande)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
		    log.Fatal(err)
		}
		fmt.Printf("test %d [%s] : ", i, table[i].Name)
		if CheckRes(outb.String(), table[i].Outpout, table[i].Type_search) {
			color.Set(color.FgGreen)
			fmt.Println(" [ OK ]")
			color.Unset()
		} else {
			color.Set(color.FgRed)
			fmt.Println(" [ KO ]")
			color.Unset()
			fmt.Println("Return : ")
			fmt.Println(outb.String())
			fmt.Printf("(type : %s [search type]), must return : \n%s\n", table[i].Type_search, table[i].Outpout)
			return
		}
	}
}

func CheckRes(Outpout string, attOuput string, Type string) (bool) {

	re := strings.Index(Outpout, attOuput)
	if re != -1 {
		return (true)
	}
	return (false)
}