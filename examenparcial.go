package main

func totalNQueens(n int) int {
	queen := make([]int, n)
	return backtrack(queen, 0)
}

func backtrack(queen []int, index int) int {
	if index == len(queen) {
		return 1
	}

	res := 0
	for i := 0; i < len(queen); i++ {
		queen[index] = i
		if check(queen, index) {
			res += backtrack(queen, index+1)
		}
	}
	return res
}

func check(queen []int, index int) bool {
	for i := 0; i < index; i++ {
		if queen[i] == queen[index] || absMinus(i, index) == absMinus(queen[i], queen[index]) {
			return false
		}
	}
	return true
}

func absMinus(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	print(totalNQueens(9))
}
