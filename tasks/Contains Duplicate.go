package main

func main() {

}

func containsDuplicate(nums []int) bool {
	unique := make(map[int]struct{})

	for _, e := range nums {
		if _, contains := unique[e]; !contains {
			unique[e] = struct{}{}
		} else {
			return true
		}
	}

	return false
}
