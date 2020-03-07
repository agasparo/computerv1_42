package input

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)


type Data struct {

	Input 	string
	Natural int
}

func GetInput(input *Data) (int) {

	args := os.Args[1:]
	length := len(args)

	if length > 1 {

        return (0)
    }
   
   	if length == 1 {

    	input.Natural = is_natural(args)
    	input.Input = args[0]
    	return (1)
    }

    ReadSTDIN(input, args)
    return (1)
}

func ReadSTDIN(input *Data, args []string) {

	tab := make([]string, 1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an Equation : ")
	text, _ := reader.ReadString('\n')
	input.Input = strings.ReplaceAll(text, "\n", "")
	tab[0] = input.Input
	input.Natural = is_natural(tab)
}

func is_natural(args []string) (int) {

	tab := strings.Split(args[0], " ")

	for i := 0; i < len(tab); i++ {

		if i % 4 == 0 {

			if tab[i + 1] != "*" {
				return (0)
			}

			if len(tab[i + 2]) < 2 {
				return (0)
			}

			if tab[i + 2][0] != 'X' || tab[i + 2][1] != '^' {
				return (0)
			}
		}
	}
	return (1)
}