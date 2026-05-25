package main

import "fmt"

/*func main() {
	fmt.Printf("请输入化学成绩：\n")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("成绩等级为：")
	if num > 90 {
		print("A")
	} else if num >= 80 {
		print("B")
	} else if num >= 60 {
		print("C")
	} else {
		print("不及格")
	}
}*/

/*func main() {
	if num := 0; num > 0 {
		fmt.Println("正数")
	} else {
		fmt.Println(num)
	}
}*/

/*
func main() {

num1 := 0
num2 := 0
op := ""4
fmt.Printf("请输入第一个数字: ")
fmt.Scanln(&num1)
fmt.Printf("请输入第二个数字: ")
fmt.Scanln(&num2)
fmt.Printf("请输入操作(+ - * /)")
fmt.Scanln(&op)
switch op {
case "+":

	fmt.Printf("%d+%d=%d\n", num1, num2, num1+num2)

case "-":

	fmt.Printf("%d-%d=%d\n", num1, num2, num1-num2)

case "*":

	fmt.Printf("%d*%d=%d\n", num1, num2, num1*num2)

case "/":

	fmt.Printf("%d/%d=%d\n", num1, num2, num1/num2)

default:

		fmt.Printf("无效操作")
	}
*/
//写法2
/*func main() {
	fmt.Printf("输入成绩：")
	s := 0
	fmt.Scanln(&s)
	switch {
	case s >= 0 && s <= 59:
		fmt.Println("不及格")
	case s <= 79:
		fmt.Println("及格")
	case s <= 89:
		fmt.Println("良好")
	case s <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("成绩有误")
	}
}*/
//写法3
/*func main() {
	year := 0
	month := 0
	day := 0
	fmt.Printf("请输入年份和月份\n")
	fmt.Scanln(&year, &month)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		{
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				day = 29
			} else {
				day = 28
			}
		}
	default:
		fmt.Printf("输入有误")
	}
	fmt.Printf("%d-%d-%d", year, month, day)
}*/
//for
/*func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
	fmt.Printf("->%d", i)
}*/
/*func main() {
	for i := 58; i >= 23; i-- {
		fmt.Printf("%d\n", i)
	}
}*/
/*func main() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Printf("res:%d\n", res)
}*/
/*func main() {
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			count++
			fmt.Printf("%2d ", i)
			if count%5 == 0 {
				fmt.Printf("\n")
			}
		}
	}
	fmt.Printf("\ncount:%d\n", count)
}*/
/*func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}*/
/*func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Print("\n")
	}
}*/
/*func main() {
out:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break out
				//continue
			}
			fmt.Printf("i=%d   j=%d\n", i, j)
		}
	}
}*/ //153
/*func main() {
	fmt.Print("请输入一个三位数：\n")
	num := 0
	fmt.Scanf("%d", &num)
	if num >= 100 && num <= 999 {
		n := num
		var res float64
		for num > 0 {
			a := num % 10
			res += math.Pow(float64(a), 3)
			num /= 10
		}
		if res == float64(n) {
			fmt.Printf("是水仙花数")
		} else {
			fmt.Printf("不是水仙花数")
		}
	} else {
		fmt.Print("输入不符合题意")
	}

}*/

/*
	func main() {
		for i := 2; i <= 100; i++ {
			flag := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag == true {
				fmt.Printf("%d\n", i)
			}
		}
	}
*/
func main() {
	n := 10
	for n < 15 {
		if n == 13 {
			goto L
		}
		fmt.Println(n)
	L:
		n++
	}

}
