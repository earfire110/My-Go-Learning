package main

import (
	"fmt"
	"math"
)

/*func main() {
	f, err := os.Open("error_test2.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}*/

/*
	func main() {
		//创建一个error数据
		err1 := errors.New("abcdef")
		fmt.Println(err1)
		fmt.Printf("%T\n", err1) //*errors.errorString
		//法2
		err2 := fmt.Errorf("ABCDEF")
		fmt.Println(err2)
		fmt.Printf("%T\n", err2)

		err3 := checkAge(-30)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		fmt.Println("go on")
	}

// 设计一个函数

	func checkAge(age int) error {
		if age < 0 {
			//return errors.New("年龄不合法")
			err := fmt.Errorf("年龄是%d 不合法", age)
			return err
		}
		fmt.Println("年龄是 ", age)
		return nil
	}
*/
/*func main() {
	addr, err := net.LookupHost("www.baidu.com")
	fmt.Println(err)
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("timeout")
		} else if ins.Temporary() {
			fmt.Println("temporary")
		} else {
			fmt.Println("unknown")
		}
	}
	fmt.Println(addr)
}*/
/*func main() {
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err) //syntax error in pattern
		return
	}
	fmt.Println(files)
}*/
//自定义错误
func main() {
	radius := -3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok {
			fmt.Printf("半径是：%.2f\n", err.radius)
		}
		return
	}
	fmt.Println("Area is", area)
}

// 1.定义结构体，表示错误类型
type areaError struct {
	msg    string //错误信息
	radius float64
}

// 2.实现error接口——实现error()方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error：半径，%.2f，%s", e.radius, e.msg)
}
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径时非法的", radius}
	}
	return math.Pi * radius * radius, nil
}
