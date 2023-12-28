package utils

import (
	"RyzeSCA/internal/constant"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	version0 = 0
	version1 = 1
	version2 = 2
)

func StrTrimSpace(v1str, v2str string) (v1, v2 string) {
	res, _ := regexp.Compile(constant.VERSIONRE)
	v1 = res.FindStringSubmatch(strings.TrimSpace(v1str))[0]
	v2 = res.FindStringSubmatch(strings.TrimSpace(v2str))[0]
	return
}
func compareSlice(v1slice, v2slice []string) int {
	for index, _ := range v1slice {
		v1, _ := strconv.Atoi(v1slice[index])
		v2, _ := strconv.Atoi(v2slice[index])
		if v1 > v2 {
			return version1
		}
		if v1 < v2 {
			return version2
		}
		if len(v1slice)-1 == index {
			return version0
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover 捕获到了异常", err)
		}
	}()
	return version0
}

// VersionComparison
/*
targetVersion > checkVersion return 1

targetVersion = checkVersion return 0

targetVersion < checkVersion return 2
*/
func VersionComparison(targetVersion string, checkVersion string) int {
	// 位数不全 后面0补齐
	targetVersion, checkVersion = StrTrimSpace(targetVersion, checkVersion)
	v1slice := strings.Split(targetVersion, ".")
	v2slice := strings.Split(checkVersion, ".")

	if len(v1slice) > len(v2slice) {
		le := len(v1slice) - len(v2slice)
		for i := 0; i < le; i++ {
			v2slice = append(v2slice, "0")
		}
	} else {
		le := len(v2slice) - len(v1slice)
		for i := 0; i < le; i++ {
			v1slice = append(v1slice, "0")
		}
	}
	slice := compareSlice(v1slice, v2slice)
	return slice
}
