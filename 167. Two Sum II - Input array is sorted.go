func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i != j && numbers[i]+numbers[j] != target {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			break
		}
	}
	return []int{i + 1, j + 1}
}
