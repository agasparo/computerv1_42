package natural

import (
	"strings"
	"fmt"
)

func Convert(str string) (string) {

	tab := strings.Split(str, " ")
	new_str := ""

	for i := 0; i < len(tab); i++ {

		if i % 4 == 0 {

			new_str += tab[i]
			new_str += " *"

			if len(tab[i + 2]) < 2 {
				new_str += " X^1"
			} else {
				new_str += " "
				new_str += tab[i + 2]
			}

			if i + 3 < len(tab) {
				new_str += " "
				new_str += tab[i + 3]
				new_str += " "
			}

			fmt.Println(new_str)
		}
	}
	return (new_str)
}