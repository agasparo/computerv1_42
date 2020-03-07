package input

import (
	"os"
	"flag"
	"agasparo"
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

	if length > 2 {
        return (0)
    }
   
   	if length == 2 {
   		if args[0][0] != '-' {
   			return (0)
   		}
    	input.Natural = is_natural(args)
    	input.Input = args[1]
    	return (1)
    }

    if args[0][0] != '-' {
   		return (0)
   	}

    ReadSTDIN(input)
    return (1)
}

func ReadSTDIN(input *Data) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an Equation : ")
	text, _ := reader.ReadString('\n')
	input.Input = strings.ReplaceAll(text, "\n", "")
}

func is_natural(args []string) (int) {

	wordPtr := flag.Int("natural", 0, "an int")
	if agasparo.Array_search(args, "-natural=1") != -1 || agasparo.Array_search(args, "-natural=0") != -1  {
		flag.Parse()
		return (*wordPtr)
	}
	return (*wordPtr)
}