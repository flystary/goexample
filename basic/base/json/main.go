package main

import "fmt"

type Person struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Weight float64 `json:"weight"`
}

//type User struct {
//	Name  	string		`json:"name"`
//	Email	string		`json:"email,omitempty"` //忽略结构体内空值字段
//	Hobby	[]string	`json:"hobby,omitempty"`
//	*Profile `json:"profile,omitempty"`   //忽略嵌套结构体内为空值的字段
//}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User              //// 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

type Card struct {
	ID    int64   `json:"id,string"`
	Score float64 `json:"score,string"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	ProFile `json:"profile"`
}

type ProFile struct {
	Hobby string `json:"hobby"`
}

func main() {
	//结构体转map
	//jsonMap()   //Go语言中的json包在序列化空接口存放的数字类型（整型、浮点型等）都会序列化成float64类型     //反射遍历结构体字段的方式生成map
	//u1 := UserInfo{Name: "q1mi", Age: 18}
	//m2, _:= ToMap(&u1, "json")
	//for k, v := range m2 {
	//	fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	//}

	//u2 := UserInfo{Name: "q1mi", Age: 18, ProFile: ProFile{"双色球"}}
	//m4,_ := ToMap2(&u2, "json")
	//for k, v := range m4 {
	//	fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	//}

	//jsonOne()
	//omitemptyDemo()
	//nestedStructDemo()
	//omitPasswordDemo()
	//inAndStringDemo()

	s1 := make([]int, 3, 10)
	var appFunc = func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}
	fmt.Println(s1)
	appFunc(s1)
	fmt.Println(s1)
	fmt.Println(s1[:5])

}
