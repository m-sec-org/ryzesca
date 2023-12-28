package utils

import "fmt"

// HandleVersionOpenClose 处理版本开闭问题
func HandleVersionOpenClose(versionEndExcluding string, versionEndIncluding string, versionStartExcluding string, versionStartIncluding string) string {
	// 开始版本 结束版本
	var startVersion string
	var endVersion string
	// 前面和后面的符号
	var startSymbol string
	var endSymbol string
	if versionStartIncluding != "" && versionStartExcluding == "" {
		startVersion = versionStartIncluding
		startSymbol = "["
	}
	if versionStartIncluding == "" && versionStartExcluding != "" {
		startVersion = versionStartExcluding
		startSymbol = "("
	}
	if versionEndIncluding != "" && versionEndExcluding == "" {
		endVersion = versionEndIncluding
		endSymbol = "]"
	}
	if versionEndIncluding == "" && versionEndExcluding != "" {
		endVersion = versionEndExcluding
		endSymbol = ")"
	}
	return fmt.Sprintf("%s%s,%s%s", startSymbol, startVersion, endVersion, endSymbol)
}
