package main

func main() {

}

type matrix struct {
	data []int
	x    int
	y    int
}

func newMatrix(x, y int) *matrix {
	return &matrix{data: make([]int, x*y), x: x, y: y}
}

func (td *matrix) Put(x, y, value int) {
	td.data[x*td.y+y] = value
}

func (td *matrix) Get(x, y int) int {
	return td.data[x*td.y+y]
}

func uniquePaths(m int, n int) int {
	arr := newMatrix(m, n)

	arr.Put(0, 0, 1)

	for i := 1; i < m; i++ {
		arr.Put(i, 0, 1)
	}

	for i := 1; i < n; i++ {
		arr.Put(0, i, 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr.Put(i, j, arr.Get(i-1, j)+arr.Get(i, j-1))
		}
	}

	return arr.Get(m-1, n-1)

}
