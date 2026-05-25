package main

import (
	"fmt"
)

/*func main() {
	var b1 bool = true
	fmt.Printf("%T,%t\n", b1, b1)
	b2 := false
	fmt.Printf("%T,%t\n", b2, b2)
	var a int
	a = 1000
	fmt.Println(a)
	var c uint8 = 100
	var d byte
	d = c
	fmt.Println(c, d)
	var e = 100
	fmt.Printf("%T,%d\n", e, e)
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.35
	fmt.Printf("%T,%.2f\n", f1, f1)
	fmt.Printf("%T,%f\n", f2, f2)
	fmt.Println(f1)
	var f3 = 1.12
	fmt.Printf("%T,%f\n", f3, f3)
}*/

/*func main() {
	var s1 string
	s1 = "asdf"
	fmt.Printf("%T,%s\n", s1, s1)
	var s2 = "hahahah"
	fmt.Printf("%T,%s\n", s2, s2)
	s3 := `hello world`
	fmt.Printf("%T,%s\n", s3, s3)
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d,%c,%q\n", v1, v1, v1, v1) //int32,65,A,'A'
	fmt.Printf("%T,%s\n", v2, v2)               //string,A

	fmt.Println("\"Hello World\"")
	fmt.Println("Hello\nWorld")
	fmt.Println(`Hel"lo Wor"ld`)
	var b bool
	fmt.Println(b)
}*/

func main() {
	var a int8
	a = 10
	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 3.84
	var c int
	c = int(f1)
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)
	//b1 := true
	//a = int8(b1)
	//fmt.Println(a, b1)
}
