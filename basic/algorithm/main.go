package main

//func main() {
//sequence() //计算斐波那契数的算法

//multiplication() //九九乘法表

//var a , b = 24, 10
//num := getMaximumCommonDivisor(a, b)
//fmt.Println("a,b的最大公约数是：", num)
//fmt.Println("a,b的最小公倍数是：", a*b/num)

//a := "1000213320001"   //回文判断数字
//isPalindrome(a)
//fmt.Println(isPalindrome(a))

//for i := 1; i < 1000; i++ {   //水仙花数
//	if isDaffodilNum(i) {
//		fmt.Println("属于水仙花数有: ", i)
//	}
//}

//var a, b time.Duration //时间重置
//a = time.Second
//b = a * 3
//fmt.Println(b)
//a = time.Duration(500) * time.Millisecond
//b = a * 1
//fmt.Println(b)

//var str []int   //求100以内的质数
//j := 0
//for i := 2; i < 100; i++ {
//	if isPrime(i) {
//		str = append(str, i)
//		j++
//	}
//}
//fmt.Println(str)

//wage() //工龄工资计算器

//x := map[string]int{  //判断两个map是否拥有相同的键和值
//	"A": 0,
//	"B": 1,
//	"C": 2,
//	"D": 3,
//}
//y := map[string]int{
//	"B": 1,
//	"C": 2,
//	"D": 3,
//}
//z := map[string]int{
//	"A": 0,
//	"B": 1,
//	"C": 2,
//	"D": 3,
//}
//fmt.Println(equal(x, y))
//fmt.Println(equal(x, z))

//student := map[string]int { //实现顺序遍历map中的元素
//	"lisa":     17,
//	"bob":      20,
//	"victoria": 24,
//	"sabit":    40,
//}
//traverseMap(student)

//factorial()  //定义一个map，存1到20的阶乘并顺序输出
//factorialSort() //阶乘计算完再顺序输出

//n := 5 //约瑟夫环问题
//k := 2
//res := josephus(n, k)
//fmt.Println(res)

//s := "A, B, 44..., 5 z ,a, b, ...,&&z   " //.给出一串字符，要求统计出里面的字母、数字、空格以及其他字符的个数。
//countUnicode(s)

//var s = []int{1, 2, 3, 4, 5, 6, 7}  //就地反转一个整型slice中的元素
//fmt.Println(reverse(s))

//x := []string{"a","b","c","d"}   //两个slice是否拥有相同的元素
//y := []string{"a","d","p","d"}
//z := []string{"a","b","c","d"}
//
//fmt.Println(equals(x, z))
//fmt.Println(equals(y, z))

//a := [][]int {  //翻转二维数组
//	{0, 1, 2, 3},
//	{4, 5, 6, 7},
//	{8, 9, 10, 11},
//}
//transpose(a)

//var i int = 10   //斐波那契数列
//var j int = 3
//fmt.Println("第", i+1, "项的值为", isFibonacciSequence(i))
//fmt.Println("第", j+1, "项的值为", isFibonacciSequence(j))
//
//
//sum := (f(10))   //猴子吃桃子
//fmt.Printf("第一天摘了%d个桃子", sum)

//readFilePath := "Concurrent/sort/artical.txt"
//writeFilePath := "Concurrent/sort/wordFrequency.txt"
//getWordFrequency(readFilePath, writeFilePath)
//
//	fmt.Println("Bubblr sorting sort in Go")
//	unsortedInput := []int{5,3,4,1,2,9,0}
//	sorted := BubbleSort(unsortedInput)
//	fmt.Println(sorted)
//
//}
