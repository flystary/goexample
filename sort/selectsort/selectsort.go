package selectsort

import "test/sort/utils"

func Sort(list []int, left, right int) {
	if left == right {
		return
	}
	minIndex := left
	for i := left; i <= right; i++ {
		if list[i] <= list[minIndex] {
			minIndex = i
		}
	}
	utils.SwapGo(list, left, minIndex)
	Sort(list, left+1, right)
}
