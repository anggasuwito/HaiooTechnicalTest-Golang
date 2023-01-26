package answer

import "fmt"

func Two(inputOne string, inputTwo string) {
	result := Calculation(inputOne, inputTwo)
	if result > 1 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}

func Calculation(inputOne string, inputTwo string) int {
	byteInputOne := []byte(inputOne)
	byteInputTwo := []byte(inputTwo)

	column := make([]int, len(byteInputOne)+1)
	for y := 1; y <= len(byteInputOne); y++ {
		column[y] = y
	}

	for x := 1; x <= len(byteInputTwo); x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= len(byteInputOne); y++ {
			oldkey := column[y]
			var incr int
			if byteInputOne[y-1] != byteInputTwo[x-1] {
				incr = 1
			}

			column[y] = Minimal(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[len(byteInputOne)]
}

func Minimal(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
