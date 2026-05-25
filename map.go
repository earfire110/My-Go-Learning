package main

import "fmt"

/*func main() {
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 87, "Java": 79}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	if map1 == nil {
		map1 = make(map[int]string)
	}
	//2.存储键值对到map中
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "hah"
	map1[7] = "xiaohong"

	//3.获取数值，根据key获取对应的value值
	//如果key存在，获取value值，如果不存在，获取value类型的零值
	fmt.Println(map1)
	fmt.Println(map1[2])
	fmt.Println(map1[30])
	fmt.Println(map1[7])
	v1, ok := map1[40]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作得key不存在，获取到的是零值：", v1)
	}
	//4.修改数据
	fmt.Println(map1)
	map1[3] = "ha"
	fmt.Println(map1)
	//5.删除数据
	delete(map1, 7)
	fmt.Println(map1)
	delete(map1, 10)
	fmt.Println(map1)
	//6.长度
	fmt.Println(len(map1))
}*/

/*
	func main() {
		m := make(map[string]int)
		m["a"] = 1
		x, ok := m["b"]
		fmt.Println(x, ok)
		x, ok = m["a"]
		fmt.Println(x, ok)
	}
*/
/*func main() {
	map1 := make(map[int]string)
	map1[1] = "h"
	map1[2] = "e"
	map1[3] = "l"
	map1[4] = "l"
	map1[5] = "o"
	//1.遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//有序遍历
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "——>", map1[i])
	}

			1.获取所有的key，存到切片/数组
		    2.进行排序
		    3.遍历

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	for k, _ := range map1 {
		keys = append(keys, k) //加入到切片中
	}
	fmt.Println(keys) //[2 3 4 5 1]
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys { //对切片操作
		fmt.Println(key, map1[key])
	}

	//s1 := []string{"apple", "cherry", "banana", "abc", "acd", "acc"}
	//fmt.Println(s1)
	//sort.Strings(s1)
	//fmt.Println(s1)
}*/
//map结合slice

/*
	func main() {
		map1 := make(map[string]string)
		map1["name"] = "小红"
		map1["age"] = "18"
		map1["sex"] = "female"
		map1["address"] = "beijing"
		fmt.Println(map1)
		map2 := map[string]string{"name": "ruby", "age": "30", "sex": "male", "address": "nanjing"}
		fmt.Println(map2)
		map3 := map[string]string{"name": "lili", "age": "25", "sex": "female", "address": "shanghai"}
		fmt.Println(map3)
		fmt.Println()
		//创建slice
		s1 := make([]map[string]string, 0, 3)
		s1 = append(s1, map1)
		s1 = append(s1, map2)
		s1 = append(s1, map3)
		for i, v := range s1 {
			fmt.Printf("第%d个人的信息是：\n", i+1)
			fmt.Printf("姓名：%s\n", v["name"])
			fmt.Printf("年龄：%s\n", v["age"])
			fmt.Printf("性别：%s\n", v["sex"])
			fmt.Printf("地址：%s\n", v["address"])
		}
	}
*/
func main() {
	map3 := make(map[string]map[string]string)
	m1 := make(map[string]string)
	m1["name"] = "lili"
	m1["age"] = "18"
	map3["first"] = m1
	m2 := make(map[string]string)
	m2["name"] = "ming"
	m2["age"] = "24"
	map3["second"] = m2
	m3 := make(map[string]string)
	m3["name"] = "hong"
	m3["age"] = "20"
	map3["third"] = m3
	fmt.Println(map3)

	map4 := make(map[string]string)
	map4["lili"] = "18"
	map4["ming"] = "20"
	map4["hong"] = "21"
	fmt.Println(map4)
	map5 := map4
	fmt.Println(map5)
	map5["lili"] = "38"
	fmt.Println(map4)
	fmt.Println(map5)

}
