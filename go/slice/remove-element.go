package slice

// 要移除第i个元素
// 使用 slice = append(slice[:i],slice[i+1:]...)

// Remove target value in slice
func RemoveElement(slices []interface{}, target interface{}) []interface{} {
	for i, s := range slices {
		if s == target {
			slices = append(slices[:i], slices[i+1:]...)
		}
	}
	return slices
}
