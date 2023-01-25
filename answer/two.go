package answer

import "fmt"

func Two(inputOne string, inputTwo string) {
	if len(inputOne)-len(inputTwo) >= 2 || len(inputTwo)-len(inputOne) >= 2 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}
