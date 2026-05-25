package main

import "fmt"

/*func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	//slice
	var s1 []int
	fmt.Println(s1)
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)
	fmt.Printf("%T ,%T\n", arr, s2)
	s3 := make([]int, 3, 5)
	fmt.Println(s3)
	fmt.Printf("长度：%d,容量：%d\n", len(s3), cap(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//append
	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1)
	fmt.Println(s4)
	s4 = append(s4, 3, 4, 5, 6, 7, 8)
	fmt.Println(s4)
	s4 = append(s4, s3...)
	fmt.Println(s4)
}*/
//在已有数组上直接创建切片
/*func main() {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8} //数组
	fmt.Println(a)
	fmt.Printf("长度：%d,容量：%d\n", len(a), cap(a))
	//创建切片
	s1 := a[:5]
	s2 := a[3:6]
	s3 := a[6:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	//fmt.Printf("%p\n", s2)
	//fmt.Printf("%p\n", s3)
	//切片的长度和容量
	fmt.Printf("长度：%d,容量：%d\n", len(s1), cap(s1))
	fmt.Printf("长度：%d,容量：%d\n", len(s2), cap(s2))
	fmt.Printf("长度：%d,容量：%d\n", len(s3), cap(s3))
	//修改数组
	a[5] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//修改切片
	s2[0] = 300
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//
	s1 = append(s1, 0, 0)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, -1, -1, -1, -1) //
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)
}*/

// 深拷贝与浅拷贝
func main() {
	//切片的深拷贝
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 0) //自动扩容
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	//copy()函数
	s3 := []int{7, 8, 9}
	//copy(s2,s3)
	//copy(s3, s2)
	copy(s3[1:], s2[3:])
	fmt.Println(s2) //[1 2 3 4 5]
	fmt.Println(s3) //[7 4 5]

}
