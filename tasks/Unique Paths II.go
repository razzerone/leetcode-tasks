package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	arr := newMatrix(m, n)

	arr.Put(0, 0, 1)

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		arr.Put(i, 0, 1)
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		arr.Put(0, i, 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			arr.Put(i, j, arr.Get(i-1, j)+arr.Get(i, j-1))
		}
	}

	return arr.Get(m-1, n-1)

}

//type matrix struct {
//	data []int
//	x    int
//	y    int
//}
//
//func newMatrix(x, y int) *matrix {
//	return &matrix{data: make([]int, x*y), x: x, y: y}
//}
//
//func (td *matrix) Put(x, y, value int) {
//	td.data[x*td.y+y] = value
//}
//
//func (td *matrix) Get(x, y int) int {
//	return td.data[x*td.y+y]
//}
