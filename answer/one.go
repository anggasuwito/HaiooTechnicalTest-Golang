package answer

import (
	"encoding/json"
	"fmt"
	"math"
)

func One(number float64) {
	number = math.Ceil(number/100) * 100
	currency := []float64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	result := make(map[string]int)
	var index int
	for {
		if index == len(currency) {
			break
		}

		if number >= currency[index] {
			number -= currency[index]
			result[fmt.Sprintf("Rp. %v", currency[index])]++
		} else {
			index++
		}
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
