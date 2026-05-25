package main

import "fmt"

/*func main() {
	const PATA string = "http:ww.baidu.com:80"
	const PI = 3.1415926
	fmt.Println(PATA)
	//fmt.Println(PI)
	const C1, C2, C3 = 100, 3.14, "hi"
	const (
		NALE   = 1
		FEMALE = 10
		UNKNOW = 3
	)
	fmt.Println(C1, C2, C3)
	fmt.Println(NALE, FEMALE, UNKNOW)
}*/
//iota
/*func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
	)
	fmt.Println(d, e)
	const (
		MANE = iota
		FEMALE
		UNKNOW
	)
	fmt.Println(MANE, FEMALE, UNKNOW)
}*/

func main() {
	const (
		A = iota
		B
		C
		D = "haha" //iota=3
		E          //iota=4
		F = 100    //iota=5
		G          //iota=6
		H = iota
		I //iota=8
	)
	const (
		J = iota
	)
	fmt.Println(A, B, C, D, E, F, G, H, I)
}
