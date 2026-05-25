package main

import (
	"fmt"
	"math"
)

// go语言中接口和类型的实现关系是非嵌入式（不需要显示声明）
// 1.当需要接口类型的对象时，可以使用任意实现类对象代替
// 2.接口对象不能访问实现类中的属性
/*func main() {
	m1 := Mouse{"hp"}
	fmt.Println(m1.name)
	m1.start()
	f1 := FlashDisk{"lenovo"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)

	var usb USB
	usb = m1
	usb.start()
	usb.end()

}

// 1.定义接口
type USB interface {
	//只写方法的声明
	start()
	end()
}

// 2.实现类
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

// 实现方法
func (m Mouse) start() {
	fmt.Println("Mouse.start")
}
func (m Mouse) end() {
	fmt.Println("Mouse.end")
}
func (f FlashDisk) start() {
	fmt.Println("FlashDisk.start")
}
func (f FlashDisk) end() {
	fmt.Println("FlashDisk.end")
}

// 3.测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
*/
/*func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you! ")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iPhone,I can call you! ")
}*/
/*func main() {
	//接口类型变量
	var a1 A = Cat{"black"}
	var a2 A = Person{"xiaoli", 18}
	var a3 A = 5.45
	var a4 A = "aaa"
	fmt.Println(a1, a2, a3, a4)
	test1(a1)
	test1(3.14)
	test2(100)
	test2("bbb")

	//map key为字符串，value任意
	map1 := make(map[string]interface{})
	map1["name"] = "xiaohua"
	map1["age"] = 23
	map1["friend"] = Person{"xiaoli", 18}
	fmt.Println(map1)
	//slice
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 200, "ccc")
	fmt.Println(slice1)
	test3(slice1)
}

// 参数设为空接口类型,函数可接受任意类型的数据
func test1(a A) {
	fmt.Println(a)
}
func test2(a interface{}) {
	fmt.Println("匿名", a)
}
func test3(slice2 []interface{}) {
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("第%d个数据是%v\n", i+1, slice2[i])
	}
}

// 空接口
type A interface {
}
type Cat struct {
	color string
}
type Person struct {
	name string
	age  int
}*/

//接口的嵌套

/*func main() {
	var a Cat = Cat{}
	a.test1()
	a.test2()
	a.test3()
	fmt.Println("--------------")
	var a2 A = a
	a2.test1()
	fmt.Println("---------------")
	var a3 C = a
	a3.test1()
	a3.test2()
	a3.test3()
	fmt.Println("---------------")

}

type A interface {
	test1()
}
type B interface {
	test2()
}
type C interface {
	A
	B
	test3()
}
type Cat struct {
}

func (c Cat) test1() {
	fmt.Println("test1")
}
func (c Cat) test2() {
	fmt.Println("test2")
}
func (c Cat) test3() {
	fmt.Println("test3")
}*/

// 接口断言
func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	testShape(s1)
	fmt.Println("-------------------------")
	getType(t1)
	getType(s1)
	var t2 *Triangle = &Triangle{30, 40, 50}
	fmt.Printf("t2:%T,%p\n", t2, &t2)
	getType(t2)
	getType2(t1)
	getType2(s1)
	getType2(t2)
}
func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T,%p\n", ins, &ins)
		fmt.Printf("s:%T,%p\n", s, &s)
	} else {
		fmt.Println("其它")
	}
}
func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形", ins.radius)
	case *Triangle:
		fmt.Println("指针三角形", ins.a, ins.b, ins.c)
	}
}
func testShape(s Shape) {
	fmt.Printf("周长：%.2f，面积：%.2f", s.peri(), s.area())
}

type Shape interface {
	peri() float64
	area() float64
}
type Triangle struct {
	a, b, c float64
}
type Circle struct {
	radius float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.peri() / 2.0
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}
func (c Circle) peri() float64 {
	return math.Pi * c.radius * 2
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
