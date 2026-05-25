package main

import "fmt"

func main() {
	fmt.Printf("输入一个整数和字符串：\n")
	var x int
	var y string
	fmt.Scanf("%d %s", &x, &y)
	fmt.Println(x, y)
	fmt.Scanln(x, y)
	fmt.Println(x+1, y)
}
