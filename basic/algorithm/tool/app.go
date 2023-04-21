package main

import "math"

//整数反转

func reserve(i int) int {
	var sum int
	for i != 0 {
		sum = sum*10 + i%10
		i /= 10
	}
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}
	return sum

}

func toLowerCase(str string) string {
	bytes := []byte(str)
	for i, v := range bytes {
		if v >= 'A' && v < 'Z' {
			bytes[i] += 'a' - 'A'
		}
	}
	return string(bytes)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	l := mergeSort(arr[0:i])
	r := mergeSort(arr[i:])
	result := merge(l, r)
	return result

}

func merge(l, r []int) []int {
	a, b := 0, 0
	c := len(l)
	d := len(r)
	result := make([]int, 0)
	for a < c && b < d {
		if l[a] > r[b] {
			result = append(result, r[b])
			b++
			continue
		}
		result = append(result, l[a])
		a++
	}
	result = append(result, r[b:]...)
	result = append(result, l[a:]...)
	return result
}
