package main

/*import "time"

func main() {
}

// 定义time.Duration的别名
type MyDuration time.Duration

func (m MyDuration) Simpleset() {
	//cannot define new methods on non-local type MyDuration
}*/

/*type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person——>", p.name)
}

type People = Person

func (p People) show2() {
	fmt.Println("Perple——>", p.name)
}

type Student struct {
	Person
	People
}

func main() {
	var s Student
	//s.name = "xiaoming"//ambiguous selector s.name
	s.Person.name = "xiaoming"
	s.Person.show()
	fmt.Printf("%T,%T\n", s.Person, s.People)
	s.People.name = "lixiaohua"
	s.People.show()  //Person——> lixiaohua
	s.People.show2() //Perple——> lixiaohua
}*/
