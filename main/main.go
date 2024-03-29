package main

import (
	"alderaan/resource"
	"fmt"
)

func AvgChange(a []float32) float32 {
	var sum float32
	for _, value := range a {
		sum += value
	}
	return sum / float32(len(a))
}

func FloatToPercentage(val float32) string {
	return fmt.Sprintf(
		"%.2f%%", val*100)
}

func mmain() {
	{
		slicee := (resource.GenRandSlice(100, 10))
		fmt.Println(slicee)
		fmt.Println("")

		res := resource.PercentageChange(slicee)
		fmt.Print(res)
		fmt.Println("")

		change := AvgChange(res)
		fmt.Print(change)
		fmt.Println("")

		fmt.Println("")
		per := FloatToPercentage(change)
		fmt.Print(per)
	}

	end := 100000
	for i := 1; i < end; i++ {
	}
}
