package natural

import (
	"strings"
	"fmt"
)

func Convert(str string) (string) {

	tab := strings.Split(str, " ")
	new_str := ""
	var x, in = 0, 0

	for i := 0; i < len(tab); i++ {

		in = 0
		x = 0
		new_str += intorx(tab[i], &x)

		if len(tab) > i + 1 && tab[i + 1] != "*" {
			new_str += " * X^"
			new_str += powerof(tab[i], x)
			in = 1
		} else if len(tab) > i + 2 {
			new_str += " * X^"
			new_str += powerof(tab[i + 2], x)
			in = 3
		} else {
			if x == 1 {
				new_str += " * X^"
				new_str += powerof(tab[i], x)
			} else {
				new_str += " * X^0"
			}
			in = 4
		}

		if i + in < len(tab) {
			new_str += " "
			new_str += tab[i + in]
			new_str += " "
		}

		i += in

		fmt.Println(new_str)
	}
	return (new_str)
}

func intorx(str string, x *int) (string) {

	if str == "X" || str[0] == 'X' {
		*x = 1
		return ("1")
	}
	return (str)
} 

func powerof(str string, x int) (string) {

	if x == 1 {
		if str[0] == 'X' {
			return (string(str[len(str) - 1]))
		}
		return ("k")
	}
	if  str[0] != 'X' {
		return ("0")
	}
	return ("1")
}