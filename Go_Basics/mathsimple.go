package main

import (
	"fmt"
)

func main() {

	/* implicit declaration of variables */
	int1, int2, int3 := 4, 2, 6
	intSum := int1 + int2 + int3
	fmt.Println("Int Sum :", intSum)

	float1, float2, float3 := 2.5, 4.5, 6.5
	floatSum := float1 + float2 + float3
	fmt.Println("Float Sum :", floatSum)

}
