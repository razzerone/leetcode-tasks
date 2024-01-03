package main

func main() {

}

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	columns := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] {
				matrix[i] = make([]int, len(matrix[0]))
				break
			}
			if columns[j] {
				matrix[i][j] = 0
			}
		}
	}
}
