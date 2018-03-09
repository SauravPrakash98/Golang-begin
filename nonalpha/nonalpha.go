package main

import "fmt"

func main() {
	var str string
	str = "a..s($4.;'1241d135"
	//	fmt.Println(str)
	var s []byte
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= '0' && str[i] <= '9') {
			s = append(s, str[i])
			//fmt.Println(i)
		} else {
			//str = strings.Trim(str, string(str[i]))
		}

	}
	st := string(s)
	fmt.Println(st)
}
