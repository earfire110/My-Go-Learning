package main

import "fmt"

/*func main() {
	w1 := Worker{"xiaoming", 30, "男"}
	w1.work()
	w2 := &Worker{"Ruby", 34, "女"}
	fmt.Printf("%T\n", w2)
	w2.work()
	w2.rest()
	fmt.Println("-----------------------------")

	w2.printInfo()
	c1 := Cat{"白色", 2}
	c1.printInfo()
}

// 1.定义一个工人结构体
type Worker struct {
	name string
	age  int
	sex  string
}
type Cat struct {
	color string
	age   int
}

// 2.定义行为方法
func (w Worker) work() {
	fmt.Println(w.name, "在工作")
}
func (p *Worker) rest() {
	fmt.Println(p.name, "在休息")
}
func (p *Worker) printInfo() {
	fmt.Printf("工人姓名：%s，年龄：%d，性别：%s\n", p.name, p.age, p.sex)
}
func (p *Cat) printInfo() {
	fmt.Printf("猫咪颜色：%s，年龄：%d\n", p.color, p.age)
}
*/
//继承中方法

func main() {
	//创建Person类型
	p1 := Person{"xiaoming", 30}
	fmt.Println(p1.name, p1.age)
	p1.eat()
	//创建Student类型
	s1 := Student{Person{"ruby", 18}, "abc"}
	fmt.Println(s1.name, s1.age) //省略Person
	fmt.Println(s1.school)

	s1.eat() //子类对象访问父类方法
	s1.eat()
	s1.Person.eat()
}

// 父类
type Person struct {
	name string
	age  int
}

// 子类
type Student struct {
	Person
	school string
}

// 方法
func (p Person) eat() {
	fmt.Println("父类方法，窝窝头")
}

func (s Student) eat() {
	fmt.Println("学生吃食堂")
}
