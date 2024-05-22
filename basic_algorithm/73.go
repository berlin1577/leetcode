package basic_algorithm

func setZeroes(matrix [][]int) {
	col, row := make([]bool, len(matrix[0])), make([]bool, len(matrix))
	for i, v := range matrix {
		for j, k := range v {
			if k == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, v := range matrix {
		for j := range v {
			if row[i] == true || col[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}
