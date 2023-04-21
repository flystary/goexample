package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func sequence() {
	const n = 40
	starttime := time.Now()
	fibN := fib(n)
	endtime := time.Now()
	costTime := endtime.Sub(starttime)
	fmt.Println(costTime)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func multiplication() {
	start := time.Now()
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
	//end := time.Now()
	tc := time.Since(start)
	//tc := end.Sub(start)
	fmt.Printf("耗时是%v", tc)
}

func getMaximumCommonDivisor(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}

func isPalindrome(a string) bool {
	var i, j int
	var b bool
	for i = 0; i < len(a)/2-1; i++ {
		j = len(a) - 1
		if a[i] != a[j] {
			b = false
		}
		b = true
		j--
	}
	return b
}
func isDaffodilNum(num int) bool {
	a := num / 100
	b := (num / 10) % 10
	c := num % 10
	result := a*a*a + b*b*b + c*c*c
	if num == result {
		return true
	}
	return false
}

func isPrime(i int) bool {
	for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func wage() {
	var n, salary, sum, a int
	var err error

	fmt.Print("请属于你的工龄：")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("请输入你的基本工资：")
	_, err = fmt.Scanln(&salary)
	if err != nil {
		fmt.Println(err)
	}
	if n >= 0 && n < 1 {
		a = 200
	} else if n >= 1 && n < 3 {
		a = 500
	} else if n >= 3 && n < 5 {
		a = 1000
	} else if n >= 5 && n < 10 {
		a = 2500
	} else if n >= 10 && n < 15 {
		a = 5000
	}
	sum = salary + a
	fmt.Printf("您目前工作了%d年，基本工资为%d元,应涨工资%d元,涨后工资%d元", n, salary, a, sum)
	return
}

/*
func mapDbWrite()  {

	//创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	//map插入对应的key - value对,各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是",countryCapitalMap[country])
	}

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	}else {
		fmt.Println("American 的首都不存在")
	}
	//用户名：密码@tcp(地址:3306)/数据库名
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.2.132:3306)/test")
	if err != nil {
		fmt.Println(err)
	}

	//往数据插入数据
	for k,v := range countryCapitalMap {
	result, err := db.Exec("INSERT INTO countryCapital(country,capital)VALUES (?,?)", k, v)
		//result, err := db.Prepare("INSERT INTO countryCapital(country,capital)VALUES ($1,$2)")
		if err != nil {
			fmt.Println(result,err)
		}
		//result, err = result.Exec(k,v)
	}



}
*/

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xValue := range x {
		if yVaalue, ok := y[k]; !ok || yVaalue != xValue {
			return false
		}

	}
	return true
}

func traverseMap(student map[string]int) {
	var names []string
	for name := range student {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s的年龄是%d\n", name, student[name])
	}
}

func factorial() {
	m := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i
		}
		fmt.Println(i, "的阶乘是", m[i])
	}

}

func factorialSort() {
	m := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i
		}
	}
	arr := make([]int, 0)
	for k, _ := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(arr[i], "的阶乘是", m[arr[i]])
	}
}

func josephus(n, k int) int {
	if n == 1 {
		return n
	}
	return (josephus(n-1, k)+k-1)%n + 1
}

func countUnicode(s string) {
	var letter int
	var num int
	var space int
	var other int

	for i := 0; i < len(s); i++ {
		if (s[i]) >= 'a' && s[i] <= 'z' || (s[i] >= 'A' && s[i] <= 'Z') {
			letter++
		} else if s[i] >= '0' && s[i] <= '9' {
			num++
		} else if s[i] == ' ' {
			space++
		} else {
			other++
		}
	}
	fmt.Printf("字母的个数为%d\n", letter)
	fmt.Printf("数字的个数为%d\n", num)
	fmt.Printf("空格的个数为%d\n", space)
	fmt.Printf("其他字符的个数为%d\n", other)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[j]
	}
	return s
}

func equals(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func transpose(a [][]int) [][]int {
	if len(a[0]) == 0 {
		return nil
	}
	row := len(a[0])
	col := len(a)
	arr := make([][]int, row, row)
	for i := 0; i <= row-1; i++ {
		arr[i] = make([]int, col)
	}
	for k := col - 1; k >= 0; k-- {
		for y := row - 1; y >= 0; y-- {
			arr[y][k] = a[k][y]
		}
	}
	fmt.Println(arr)
	return arr
}

func isFibonacciSequence(i int) int {
	if i <= 1 {
		return 1
	}
	return isFibonacciSequence(i-1) + isFibonacciSequence(i-2)
}

func f(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2*f(n-1) + 2
	}
}

func sellGoldfish() {
	res := 11
	for j := 4; j >= 1; j-- {
		res = (res*(j+1) + 1) / j
	}

	fmt.Println("原来的鱼缸中共有", res, "条鱼")
}

func getWordFrequency(readFilePath string, writeFilePath string) {
	var fileText string
	var wordFrequencyMap = make(map[string]int)

	fileData, err := ioutil.ReadFile(readFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fileText = string(fileData)
	f := func(c rune) bool {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			return true
		}
		return false
	}
	arr := strings.FieldsFunc(fileText, f)
	for _, v := range arr {
		if _, ok := wordFrequencyMap[v]; ok {
			wordFrequencyMap[v] = wordFrequencyMap[v] + 1
		} else {
			wordFrequencyMap[v] = 1
		}
	}
	type wordFrequencyNum struct {
		Word string
		Num  int
	}
	var lstWordFrequencyNum []wordFrequencyNum
	for k, v := range wordFrequencyMap {
		lstWordFrequencyNum = append(lstWordFrequencyNum, wordFrequencyNum{k, v})
	}
	sort.Slice(lstWordFrequencyNum, func(i, j int) bool {
		return lstWordFrequencyNum[i].Num > lstWordFrequencyNum[j].Num
	})
	fmt.Println("安装单词出现的频率由高到底排序", lstWordFrequencyNum)
	////写入文件
	var arrJsonBytes string
	arrJsonBytes = arrJsonBytes + "[" + "\r\n"
	for k, v := range lstWordFrequencyNum {

		arrJsonBytes = arrJsonBytes + "{" + "\"" + v.Word + "\"" + ":" + strconv.Itoa(v.Num) + "}"
		if k != len(lstWordFrequencyNum)-1 {
			arrJsonBytes = arrJsonBytes + "," + "\r\n"
		}

	}
	arrJsonBytes = arrJsonBytes + "\r\n" + "]"
	err = ioutil.WriteFile(writeFilePath, []byte(arrJsonBytes), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	/*
		var jsonBytes []byte
		var arrJsonBytes string
		for _, v := range lstWordFrequencyNum {
			jsonBytes, err = json.Marshal(v)
			if err != nil {
				log.Fatal(err)
			}
			arrJsonBytes = arrJsonBytes + string(jsonBytes)
		}
		err = ioutil.WriteFile(writeFilePath, []byte(arrJsonBytes), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

	*/
}

func BubbleSort(input []int) []int {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				fmt.Println("Swapping")
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

//func QuickSort(array []int) []int  {
//	if len(array) < 2 {
//		return array
//	}
//
//	// we want the left-most and the right-most index of
//	// the array we are going to sort
//	left, right := 0, len(array)-1
//
//	// choose a random pivot
//	pivotIndex := rand.Int() % len(array)
//	array[pivotIndex], array[right] = array[right], array[pivotIndex]
//
//	for i := range array {
//		if array[i] < array[right] {
//			array[i], array[left] = array[left],array[i]
//			left++
//		}
//	}
//	array[left],array[right] = array[right], array[left]
//	QuickSort(array[:left])
//	QuickSort(array[left+1:])
//
//	return array
//}

func Sort(list []int, left, right int) {
	if right < left {
		return
	}
	flag := list[left]
	start := left
	end := right
	for {
		if start == end {
			break
		}
		for list[end] >= flag && end > start {
			end--
		}
		for list[start] <= flag && end > start {
			start++
		}
		if end > start {
			SwapGo(list, start, end)
		}
	}
	SwapGo(list, left, start)
	Sort(list, left, start-1)
	Sort(list, start+1, right)

}
