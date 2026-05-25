package main

import "fmt"

func main() {
	funA()
	defer myprint("defer main:3") //执行
	funB()
	defer myprint("defer main:4")
}
func myprint(s string) {
	fmt.Println(s)
}
func funA() {
	fmt.Println("funA")
}
func funB() {
	defer func() {
		//recover不需要参数，有返回值，是当前程序panic引发恐慌时所传递的参数
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序被恢复")
		}
	}()
	fmt.Println("funB")
	defer fmt.Printf("defer fun():1\n") //执行
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("funB函数恐慌") //函数发生恐慌已经被defer的函数也会执行
		}
	}
	defer myprint("defer fun():2")
}
