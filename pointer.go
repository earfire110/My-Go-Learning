package main

/*func main() {
	a := 10
	fmt.Println(a)         //10
	fmt.Printf("%p\n", &a) //0x14598200a0
	var p *int
	p = &a
	fmt.Println(p)         //0x14598200a0
	fmt.Printf("%p\n", &p) //0x145981a068
	fmt.Println(*p)        //10

	q := &p
	fmt.Println(q)         //0x145981a068
	fmt.Printf("%p\n", &q) //0x145981a070
	fmt.Println(*q)        //0x14598200a0
	fmt.Println(**q)       //10
}*/
//数组指针和指针数组
/*func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	//数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[1 2 3 4] 表示数组指针
	fmt.Printf("%p\n", p1)  //数组arr1的地址
	fmt.Printf("%p\n", &p1) //p1指针的地址
	//根据数组指针，操作数组
	(*p1)[0] = 10
	fmt.Println(arr1)
	p1[0] = 100 //简化写法
	fmt.Println(arr1)

	//指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}      //[1 2 3 4]
	arr3 := [4]*int{&a, &b, &c, &d} //[0x2fd07ca080f0 0x2fd07ca080f8 0x2fd07ca08100 0x2fd07ca08108]
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2) //[100 2 3 4]
	fmt.Println(a)    //1
	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a) //200
	b = 1000
	fmt.Println(arr2) //[100 2 3 4]值传递
	fmt.Println(arr3) //地址未变
	for i := 0; i < 4; i++ {
		fmt.Println(*arr3[i])
	} //200 1000 3 4 值变化
}*/
/*func main() {
	var a func()
	a = fun1
	a() //fun1()

	arr1 := fun2()                              //值传递，传递后arr被销毁数组不存在
	fmt.Printf("%T，%p，%v\n", arr1, &arr1, arr1) //[4]int，0x16e71d9b81e0，[1 2 3 4]
	arr2 := fun3()                              //数组指针//fun3结束之后。arr被销毁，引用断开，数组还存在
	fmt.Printf("%T，%p，%v\n", arr2, &arr2, arr2) //*[4]int，0x16e71d9aa068，&[5 6 7 8]
	fmt.Printf("arr2指针中存储的数组的地址：%p\n", arr2)    //0x16e71d9b8260
}
func fun1() {
	fmt.Println("fun1()")
}
func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}
func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("arr的地址：%p\n", &arr) //0x16e71d9b8260
	return &arr
}*/
