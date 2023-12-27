package utils

// SliceFindString 工具判断字符是否在切片中
// 在切片中返回true
// 不在切片中返回false
func SliceFindString(slice []string, val string) bool {
	flag := false
	for _, item := range slice {
		if item == val {
			flag = true
		}
	}
	return flag
}
