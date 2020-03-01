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

	AllTests := make([]Testes, 1, 100)
	AddTest(AllTests, "syntax error 1 (double egale)", "/Users/arthur/Desktop/computorv1/main.go", "2 * X^0 == 1 * X^1", "Notice", "regex")
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

func AddTest(table []Testes, name string, filename string, Commande string, Outpout string, Type string) {

	NewTest := Testes{
		name,
		filename,
		Commande,
		Outpout,
		Type,
	}
	table[len(table) - 1] = NewTest
}

func Run(table []Testes) {

	for i := 0; i < len(table); i++ {

		cmd := exec.Command("go", "run", table[i].FileName, table[i].Commande)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
		    log.Fatal(err)
		}
		fmt.Printf("test [%s] : ", table[i].Name)
		if CheckRes(outb.String(), table[i].Outpout, table[i].Type_search) {
			color.Set(color.FgGreen)
			fmt.Println(" [ OK ]")
			color.Unset()
		} else {
			color.Set(color.FgRed)
			fmt.Println(" [ KO ]")
			color.Unset()
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