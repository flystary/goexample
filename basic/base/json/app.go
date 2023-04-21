package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//结构体转map
func jsonMap() {
	u1 := UserInfo{
		Name: "q1mi",
		Age:  18,
	}
	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		//fmt.Printf("key:%v value:%v\n", k, v)
		//key:name value:q1mi
		//key:age value:18

		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
		//key:name value:q1mi value type:string
		//key:age value:18 value type:float64
	}
}

func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func ToMap2(in interface{}, tag string) (map[string]interface{}, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	out := make(map[string]interface{})
	queue := make([]interface{}, 0, 1)
	queue = append(queue, in)

	for len(queue) > 0 {
		v := reflect.ValueOf(queue[0])
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		queue = queue[1:]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			if vi.Kind() == reflect.Ptr {
				vi = vi.Elem()
				if vi.Kind() == reflect.Struct {
					queue = append(queue, vi.Interface())
				} else {
					ti := t.Field(i)
					if tagValue := ti.Tag.Get(tag); tagValue != "" {
						out[tagValue] = vi.Interface()
					}
				}
				break
			}
			if vi.Kind() == reflect.Struct {
				queue = append(queue, vi.Interface())
				break
			}
			ti := t.Field(i)
			if tagValue := ti.Tag.Get(tag); tagValue != "" {
				out[tagValue] = vi.Interface()
			}
		}
	}
	return out, nil
}

func jsonOne() {
	p1 := Person{
		Name:   "Tim",
		Age:    18,
		Weight: 31.4,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

func omitemptyDemo() {
	u1 := User{
		Name: "hello",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}

//func nestedStructDemo() {
//	u1 := User{
//		Name: "花园",
//		Hobby: []string{"足球","篮球","双色球"},
//	}
//	b, err := json.Marshal(u1)
//	if err != nil {
//		fmt.Printf("json.Marshal failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("str:%s\n", b)
//}

func omitPasswordDemo() {
	u1 := User{
		Name:     "李权霖",
		Password: "qweasdRF123",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}

func inAndStringDemo() {
	jsonStr1 := `{"id":"1234567","score":"88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}
