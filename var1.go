package main

import "fmt"

func main() {
	var num int = 100
	fmt.Printf("num的数值是%d，地址%p\n", num, &num)
	num = 200
	fmt.Printf("num的数值是%d，地址%p\n", num, &num)

}

/*func main() {
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是：%d\n", num1)
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)
	var name = "geng"
	fmt.Printf("类型是：%T，数值是%s\n", name, name)
	var number = 70
	fmt.Printf("类型是：%T，数值是%d\n", number, number)
	sum := 100
	fmt.Println(sum)
	var a, b, c int
	fmt.Println(a, b, c)
	var m, n int = 1, 2
	fmt.Println(m, n)
	var d, e, f = 11, 12, "go"
	fmt.Println(d, e, f)
	var (
		Name = "小明"
		age  = 18
		sex  = "男"
	)
	fmt.Println(Name, age, sex)
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", Name, age, sex)

}*/
