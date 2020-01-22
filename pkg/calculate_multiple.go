package pkg

import (
	"math"
	"strconv"
)

func convertToNineBin(n int32) int64 {
	var temp int64
	var tmp = n
	var pos = 0
	for tmp > 0 {
		temp += (int64(tmp) % 2) * int64(math.Pow(10, float64(pos)))
		tmp = tmp / 2
		pos++
	}
	return temp*9
}

func GetMultiple(n int32) string {
	var valid = false
	var tmp int32 = 1
	for !valid {
		var rem = convertToNineBin(tmp) % int64(n)
		if rem != 0 {
			tmp += 1
		} else {
			valid = true
		}
	}
	result := strconv.FormatInt(convertToNineBin(tmp), 10)
	return result
}
