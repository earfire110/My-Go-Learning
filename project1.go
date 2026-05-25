package main

import "fmt"

func main() {
	fmt.Printf("%T:%v\n", 2.71828, 2.71828)
	fmt.Printf("%T:%v\n", 1.e+0, 1.e+0)
	fmt.Printf("%T:%v\n", 6.67428e-11, 6.67428e-11)
	fmt.Printf("%T:%v\n", 1e6, 1e6)
	fmt.Printf("%T:%v\n", .25, .25)
	fmt.Printf("%T:%v\n", .12345e+5, .12345e+5)

}
