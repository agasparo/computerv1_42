package natural

import (
	"strings"
	"fmt"
)

func Convert(str string) (string) {

	tab := strings.Split(str, " ")
	new_str := ""
	var in, x = 0, 0

	for i := 0; i < len(tab); i++ {

		in = 0
		x = 0
		new_str += intorx(tab[i], &x)
		new_str += " *"

		if i + 2 < len(tab) && len(tab[i + 2]) < 3 {
			new_str += " X^"
			new_str += powerof(tab[i + 2], x)
			if tab[i + 1] == "*" {
				in = 3
			} else {
				in = 1
			}
		} else if i + 2 < len(tab) {
			new_str += " "
			new_str += tab[i + 2]
			if x == 1 {
				in = 1
			}
		}

		if i + in < len(tab) {
			new_str += " "
			new_str += tab[i + in]
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

func intorx(str string, x *int) (string) {

	if str == "X" || str[0] == 'X' {
		*x = 1
		return ("1")
	}
	return (str)
} 

func powerof(str string, x int) (string) {

	if x == 1 {
		return ("1")
	}
	if  str[0] != 'X' {
		return ("0")
	}
	return ("1")
}