package main

import (
	"fmt"
	"github.com/mhcassiem/SpecialMultiple/pkg"
)

func main()  {

	inputs := []int32{23, 262, 434, 469, 450, 217, 365, 442, 456, 292, 425, 347, 496, 461, 453, 188, 179, 191, 342, 104, 246, 222, 122, 481}

	for index, value := range inputs {
		data := pkg.GetMultiple(value)
		fmt.Printf("index %d, value: %d result: %d \n", index, value, data)
	}
	//var value int32 = 5831
	//data := pkg.GetMultiple(value)
	//fmt.Printf("Result: ", data)
}
