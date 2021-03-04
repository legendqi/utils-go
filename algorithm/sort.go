/* coding: utf-8
@Time :   2021/3/4 下午5:07
@Author : legend
@File :   sort.go
*/
package algorithm

/*
快速排序
left是开始排序的index
right是结束排序的index
*/
func QuickSort(nums []int, left, right int) {
	val := nums[(left+right)/2]
	i, j := left, right
	for nums[j] > val {
		j--
	}
	for nums[i] < val {
		i++
	}
	nums[i], nums[j] = nums[j], nums[i]
	i++
	j--
	if i < right {
		QuickSort(nums, i, right)
	}
	if j > left {
		QuickSort(nums, left, j)
	}
}

/*
冒泡排序
*/
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

//二分查找法
func BinarySearch(arr *[]int, leftIndex int, rightIndex int, findValue int) int {
	if leftIndex > rightIndex {
		return -1
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		return BinarySearch(arr, leftIndex, middle-1, findValue)
	} else if (*arr)[middle] < findValue {
		return BinarySearch(arr, middle+1, rightIndex, findValue)
	} else {
		return middle
	}
}

//斐波拉契数列
func FibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
	}
}

//迭代法
func FibonacciFind(n int) int {
	x, y, fib := 0, 1, 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			fib = 0
		} else if i == 1 {
			fib = x + y
		} else {
			fib = x + y
			x, y = y, fib
		}
	}
	return fib
}
