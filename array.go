package main

import "fmt"

/*
	func main() {
		var a [4]int ////无定义自动补0
		fmt.Println(a)
		var b = [4]int{1, 2, 3}
		fmt.Println(b)
		var c = [5]int{1: 1, 3: 2}
		fmt.Println(c)
		var e = [5]string{"rose", "王二狗", "ruby"}
		fmt.Println(e)
		var d = [5]int{'A', 'B', 'C', 'D'}
		fmt.Println(d)
		f := [...]int{1, 2, 3, 4, 5}
		fmt.Println(f)
		g := [...]int{1: 3, 8: 7}
		fmt.Println(g)
		fmt.Println(len(g))
	}
*/
/*func main() {
/*var arr [5]int
fmt.Printf("%p\n", &arr)
fmt.Printf("%p\n", &arr[0])
fmt.Printf("%p\n", &arr[1])
fmt.Printf("%p\n", &arr[2])
fmt.Printf("%p\n", &arr[3])
fmt.Printf("%p\n", &arr[4])*/
/*arr1 := [5]int{1, 2, 3, 4, 5}
for index, value := range arr1 {
	fmt.Printf("下标是：%d,数值是:%d\n", index, value)
}
sum := 0
for _, v := range arr1 {
	sum += v
}
fmt.Println(sum)*/
/*var arr = [5]int{23, 56, 34, 10, 9}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	fmt.Println(arr)
}*/

// 选择排序
/*func main() {
	var arr = [5]int{23, 56, 34, 10, 9}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(arr)
	}
}*/
/*func main() {

	var arr = []int{23, 56, 34, 10, 9}
	sort.Ints(arr)
	fmt.Println(arr)

}*/
//遍历二维数组
/*func main() {
	var arr1 [3][4]int
	arr1 = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%4d", arr1[i][j])
		}
		fmt.Println()
	}

	var arr2 = [3][4]int{{3, 5, 12, 78}, {45, 62, 12, 90}, {23, 5, 15, 88}}
	for _, arr := range arr2 {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}*/
//每次都取下标和对应的数值
func main() {
	a := [...]float64{67.7, 3.14, 32, 4}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println(sum)
}
