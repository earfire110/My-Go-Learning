package main

// 结构体的定义
/*func main() {
	//1
	var s1 Person
	//fmt.Println(s1)
	s1.name = "小红"
	s1.age = 18
	s1.sex = "female"
	fmt.Println(s1)
	//2
	s2 := Person{}
	s2.name = "xiaoming"
	s2.age = 19
	s2.sex = "male"
	fmt.Println(s2)
	//3
	s3 := Person{name: "xiaoli", sex: "female", age: 20} //不用考虑顺序
	fmt.Println(s3)
	//4
	s4 := Person{"abc", 21, "male"}
	fmt.Println(s4)
}

type Person struct {
	name string
	age  int
	sex  string
}*/

// 结构体指针
/*unc main() {
	//fmt.Printf("%T\n",s1)
	s1 := Person{"xiaoming", 19, 90}
	fmt.Printf("%p,%T\n", &s1, s1)
	//var p1 *Person
	p1 := &s1
	fmt.Printf("%T\n", p1)
	fmt.Printf("%p,%T\n", &p1, p1)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println("-----------------------")
	p2 := new(Person)
	fmt.Printf("%T\n", p2)
	p2.name = "allll"
	p2.age = 24
	p2.score = 78.9
	fmt.Println(p2)
	fmt.Println(*p2)
}

type Person struct {
	name  string
	age   int
	score float64
}*/

// 匿名字段
/*func main() {
	var s1 Student
	s1 = Student{"xiaohong", 18}
	fmt.Println(s1)
	s2 := struct {
		name string
		age  int
	}{
		"xiaoli",
		30,
	}
	fmt.Println(s2)

}

type Student struct {
	name string
	age int
}*/
//结构体的嵌套

/*type Student struct {
	name string
	age  int
	book Book
}
type Book struct {
	title string
	price float64
}
*/
