package main

// TwoSum return the two indeice of arr, the elements added equ target
func TwoSum(arr []int, target int) (i, j int, ok bool) {
	m := make(map[int]int)
	for k, v := range arr {
		diff := target - v
		if dk, ok := m[diff]; ok {
			return dk, k, true
		}
		// fmt.Println(m)
		m[v] = k
	}

	return
}
