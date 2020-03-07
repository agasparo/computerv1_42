package natural

import (
	"strings"
	"fmt"
)

func Convert(str string) (string) {

	tab := strings.Split(str, " ")
	new_str := ""
	var in = 0

	for i := 0; i < len(tab); i++ {

		in = 0
		new_str += intorx(tab[i])
		new_str += " *"

		if len(tab[i + 2]) < 3 {
			new_str += " X^"
			new_str += powerof(tab[i + 2])
			if tab[i + 1] == "*" {
				in = 3
			} else {
				in = 1
			}
		} else {
			new_str += " "
			new_str += tab[i + 2]
		}

		if i + 3 < len(tab) {
			new_str += " "
			new_str += "|"//tab[i + 3]
			new_str += " "
		}

		if in > 0 {
			i += in
		} else {
			i += 3
		}
		fmt.Println(new_str)
	}
	return (new_str)
}

func intorx(str string) (string) {

	if str == "X" {
		return ("1")
	}
	return (str)
} 

func powerof(str string) (string) {

	if  str[0] != 'X' {
		return ("0")
	}
	return ("1")
}