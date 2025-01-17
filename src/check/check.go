package check

import (
	"strconv"
	"strings"
)

type CheckParam struct {
	
	Tab_str 	[]string
	Error 		string
	Error_value int
	Error_desc 	string
}

type CheckPuis struct {

	Puis0 float64
	Puis1 float64
	Puis2 float64
}

func SetErrors(input *CheckParam, error string, error_value int) {

	input.Error = error
	input.Error_value = error_value + 2
	if error_value > 0 {
		input.Error_desc = input.Tab_str[error_value]
	} else {
		input.Error_desc = input.Tab_str[0]
	}
}

func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}

func ToInt(s string) (int) {

	if i, err := strconv.Atoi(s); err == nil {
    	return (i)
	}
	return (-1)
}

func SetPuis(input *CheckParam, eqs *CheckPuis, deb int) (re int) {

	for i := deb; i < len(input.Tab_str) && input.Tab_str[i] != "="; i++ {
		
		if input.Tab_str[i][0] != 'X' && len(input.Tab_str[i]) > 1 &&input.Tab_str[i][1] == '^' {

			SetErrors(input, "You must have an X befor ^", i)
			return
		}

		if input.Tab_str[i][0] == 'X' {
			
			if i - 1 < 0 || input.Tab_str[i - 1] != "*" {
				SetErrors(input, "You must have * between number and power", i - 1)
				return (0)
			}

			if input.Tab_str[i][1] != '^' {
				SetErrors(input, "You must have ^ between X and power", i)
				return (0)
			}

			if i - 2 < 0 || !IsNumeric(input.Tab_str[i - 2]) {
				in := 0
				if !IsNumeric(input.Tab_str[i - 2]) {
					if strings.Index(input.Tab_str[i - 2], "/") != -1 {
						new_str := strings.ReplaceAll(input.Tab_str[i - 2], "/", "")
						if !IsNumeric(new_str) {
							SetErrors(input, "You must have a number before the power", i - 2)
							return (0)
						} else {
							ex := strings.Split(input.Tab_str[i - 2], "/")
							t2, _ := strconv.ParseFloat(ex[1], 64)
							if t2 == 0 {
								SetErrors(input, "no division by 0", i - 2)
								return (0)
							}
							in = 1
						}
					}
				}
				if in != 1 {
					SetErrors(input, "You must have a number before the power", i - 2)
					return (0)
				}
			}

			terme := strings.ReplaceAll(input.Tab_str[i], "X^", "")
			if !IsNumeric(terme) {
				SetErrors(input, "You must have a number for the power", i)
				return (0)
			}

			terme_int := ToInt(terme)
			if terme_int > 2 || terme_int < 0 {
				SetErrors(input, "The power must be an int between 0 and 2", i)
				return (0)
			}

			if i - 3 >= 0 {
				InitPow(input.Tab_str[i - 2], input.Tab_str[i - 3], eqs, terme_int)
			} else {
				InitPow(input.Tab_str[i - 2], "+", eqs, terme_int)
			}
		}
	}
	return (1)
}

func InitPow(s string, sign string, eqs *CheckPuis, pow int) {

	var f float64

	if strings.Index(s, "/") != -1 {
		ex := strings.Split(s, "/")
		t1, _ := strconv.ParseFloat(ex[0], 64)
		t2, _ := strconv.ParseFloat(ex[1], 64)
		if t2 == 0 {
			return
		}
		f = t1 / t2
	} else {
		f, _ = strconv.ParseFloat(s, 64)
	}
	if sign == "-" {
		f *= -1 
	}
	if pow == 2 {
		eqs.Puis2 += f 
	} else if pow == 1 {
		eqs.Puis1 += f 
	} else if pow == 0 {
		eqs.Puis0 += f 
	}
}