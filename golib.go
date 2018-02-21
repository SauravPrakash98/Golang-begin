package main

import (
	"fmt"
)

func main() {
	str := "123"
	//for i := 0; i < len(str); i++ {
	//	if str[i] == '.' {
	//			num := strconv.Atoi(str)
	//			fmt.Print(num)
	//		}
	//	}
	num := 0
	for i := len(str) - 1; i >= 0; i-- {
		num = str[i] * Pow(10, len(str)-i-1)
	}
	fmt.Print(str)
}

/*	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
        str = str[0:i]
				num := strconv.Atoi(str)
				fmt.Print(num)
			}
		} */
/*	num := 0
	for i := len(str) - 1; i >= 0; i-- {
		var pow float64
		pow = len(str) - i - 1
		num = (float64)(str[i]-'0') * math.Pow(10, pow)
	} */

/*  for i := 0; i < len(str); i++ {
		if str[i] <= 'z' && str[i] >= 'A' && str[i] >= 'a' && str[i] <= 'z' && str[i] <= '9' && str[i] >= '0' {
			var a string
			a = (string)str[i]
			t := strings.Trim(str, a)
		}
		fmt.Printf("%v", str[i])
	}
	fmt.Printf(str)
} */
