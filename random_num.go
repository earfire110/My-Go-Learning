package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(rand.Intn(10))
	//}
	str1 := rand.Float64()
	fmt.Println(str1)
	str2 := rand.ExpFloat64()
	fmt.Println(str2)
	str3 := rand.Float32()
	fmt.Println(str3)
}
