package main

import "fmt"

/*func main() {
	a := 3
	b := 2
	c := 5
	res5 := a > b && c%a == b && a < (c/b) //1 0 0
	fmt.Println(res5)
	res6 := a > b || c%a == b || a < (c/b)
	fmt.Println(res6)
}*/

func main() {
	a := 60 //0011 1100
	b := 13 //0000 1101
	//&    0000 1100
	//|    0011 1101
	//^    0011 0001
	//a&^b    0011 0000
	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)
	fmt.Printf("%b %d\n", a&b, a&b)
	fmt.Printf("%b %d\n", a|b, a|b)
	fmt.Printf("%b %d\n", a^b, a^b)
	fmt.Printf("%b %d\n", a&^b, a&^b)
	fmt.Printf("%b %d\n", ^a, ^a)
	c := 8
	fmt.Println(c << 2)
}
