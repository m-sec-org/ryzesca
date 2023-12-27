package utils

import "fmt"

// HandleVersionOpenClose 处理版本开闭问题
func HandleVersionOpenClose(versionEndExcluding string, versionEndIncluding string, versionStartExcluding string, versionStartIncluding string) string {
	// 分两个版本信息一个是开始版本一个是结束版本
	var startVersion string
	var endVersion string
	// 还有两个变量是用来表示前面和后面的符号
	var startSymbol string
	var endSymbol string
	//if versionStartIncluding != "" && versionStartExcluding != "" {
	//	startVersion = ""
	//	startSymbol = "("
	//}
	//if versionEndIncluding != "" && versionEndExcluding != "" {
	//	endVersion = ""
	//	endSymbol = ")"
	//}
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
