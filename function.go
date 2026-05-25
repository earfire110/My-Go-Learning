package main

import "fmt"

/*func getSum() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func addNum(a, b int) int {
	return a - b
}

func main() {
	//fmt.Println("aaa")
	getSum()
	fmt.Println(addNum(1, 2))
}*/

/*
	func main() {
		getSum(3, 67, 10)
		//切片
		s1 := []int{1, 2, 3, 4, 5}
		getSum(s1...)
	}

	func getSum(sums ...int) {
		//fmt.Printf("%T\n", sums) //[]int//切片其实就是长度可变的数组
		sum := 0
		for _, n := range sums {
			fmt.Printf("加入数字 %d\n", n)
			sum += n
		}
		fmt.Println(sum)
		for i := 0; i < len(sums); i++ {
			sum += sums[i]
		}
		fmt.Println(sum)
	}
*/
/*func main() {
	fmt.Println(addNum1(2, 3))
	fmt.Println(addNum2(1, 3))
}

func addNum1(x, y int) int {
	sum := x + y
	return sum
}
func addNum2(x, y int) (sum int) {
	sum = x + y
	return
}*/
//求一个矩形的周长和面积
/*func main() {
	fmt.Printf("请分别输入矩形的长和宽：")
	var a, b float64
	// fmt.Scanf("%f %f", &a, &b)
	fmt.Scanln(&a, &b)
	fmt.Printf("周长是：%.2f\n", zhouChang(a, b))
	fmt.Printf("面积是：%.2f\n", mianJi(a, b))
}

func zhouChang(x, y float64) (l float64) {
	l = float64(2) * (x + y)
	return
}
func mianJi(x, y float64) float64 {
	s := x * y
	return s
}*/
/*func main() {
	res1, res2 := rectangle(5, 3)
	fmt.Println("周长：", res1, "面积：", res2)
	res3, res4 := rectangle(2, 6)
	fmt.Println("周长：", res3, "面积：", res4)
}
func rectangle(l, w float64) (float64, float64) {
	perimeter := 2 * (l + w)
	area := l * w
	return perimeter, area
}
func rectangle2(l, w float64) (perimeter float64, area float64) {
	perimeter = 2 * (l + w)
	area = l * w
	return
}*/

/*func main() {
	n := 100
	fmt.Println(n)
	if a := 1; a < 10 {
		n := 20
		fmt.Println(n)
	}
	fmt.Println(n)
}*/
//递归函数

/*func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
*/

/*func main() {
	fun1("hello")
	defer fmt.Println("world")
	defer fun1("1234")
	fmt.Println("haha")
}
func fun1(s string) {
	fmt.Println(s)
}*/

/*func main() {
	a := 2
	fmt.Println(a)
	defer fun2(a)
	a++
	fmt.Println("main中", a)
}
func fun2(a int) {
	fmt.Println("fun2()函数中打印a：", a)
}*/

/*func main() {
	fmt.Printf("%T\n", fun1) //fun1后无括号，表示函数本身而非调用
	fmt.Printf("%T\n", fun2)
}
func fun1() {}
func fun2(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}*/
//本质
/*func main() {
	//1.函数做一个变量
	fmt.Printf("%T\n", fun1) //func(int, int)
	fmt.Printf("%p\n", fun1) //0x7ff6237561e0
	//函数名所对应的函数体的地址

	//2.直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c) //<nil>
	c = fun1
	fmt.Println(c) //0x7ff7a33464c0
	fun1(10, 20)   //a:10,b:20
	c(100, 200)    //a:100,b:200

	//fmt.Println(fun2)
	res1 := fun2            //将fun2的值(函数的地址)赋值给res1，res1和fun2指向同一个函数体
	res2 := fun2(1, 2)      //将fun2函数进行调用，将函数的执行结果赋值给res2，相当于a+b
	fmt.Println(res1, res2) //0x7ff72ddf6620 3
	fmt.Println(res1(3, 10))

}
func fun1(a, b int) {
	fmt.Printf("a:%d,b:%d\n", a, b)
}
func fun2(a, b int) int {
	return a + b
}*/

// 匿名函数
/*func main() {
	//1.匿名函数
	func() {
		fmt.Println("Hello World")
	}()
	fun1 := func() {
		fmt.Println("hello world")
	}
	fun1()
	fun1()

	//2.定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//3.定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数调用了，经执行结果给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值(地址)赋值给res2
	fmt.Println(res2)
	fmt.Println(res2(100, 200))
}*/

/*func main() {
	res1 := add(1, 2)
	fmt.Println(res1)
	res2 := oper(10, 20, add) //10 20 0x7ff6e5a74ae0
	fmt.Println(res2)         //30

	fun1 := func(a, b int) int {
		return a * b
	}
	res3 := oper(3, 4, fun1) //10 20 0x7ff6e5a74ae0
	fmt.Println(res3)

	res4 := oper(30, 6, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(res4)
}
func add(a, b int) int {
	return a + b
}
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b) //add函数在此执行
	return res
}*/

func main() {
	res1 := increment()      //res1=fun
	fmt.Printf("%T\n", res1) //func() int
	fmt.Println(res1)        //0x7ff72ebb64a0
	v1 := res1()
	fmt.Println(v1) //1
	v2 := res1()
	fmt.Println(v2)     //2
	fmt.Println(res1()) //3
	fmt.Println(res1()) //4
	fmt.Println(res1()) //5

	//每次调用increment都会产生对应的变量i，每当内层函数被调用时，相应的i进行变化
	res2 := increment()
	fmt.Println(res2) //0x7ff65b6e6800  新的地址
	v3 := res2()
	fmt.Println(v3)     //1
	fmt.Println(res2()) //2

	fmt.Println(res1()) //6
}
func increment() func() int { //外层函数
	//1.定义一个局部变量
	i := 0
	//2.定义一个匿名函数
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回匿名函数//直接return内层函数或者将内层函数赋值给一个变量
	return fun
}
