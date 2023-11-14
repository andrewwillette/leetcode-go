package luckynumbers

func luckyNumbers(matrix [][]int) []int {
	res := []int{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			rowMin := getMinimumFromRow(row, matrix)
			colMax := getMaximumFromCol(col, matrix)
			if rowMin == matrix[row][col] && colMax == matrix[row][col] {
				res = append(res, matrix[row][col])
			}
		}
	}
	return res
}

func getMinimumFromRow(row int, matrix [][]int) int {
	min := -1
	for i := 0; i < len(matrix[row]); i++ {
		if min == -1 {
			min = matrix[row][i]
		}
		if matrix[row][i] < min {
			min = matrix[row][i]
		}
	}
	return min
}

func getMaximumFromCol(col int, matrix [][]int) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] > max {
			max = matrix[i][col]
		}
	}
	return max
}
