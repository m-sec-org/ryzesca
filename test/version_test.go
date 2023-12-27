package test

import (
	"fmt"
	"strings"
	"testing"
)

var (
	version0 = 0
	version1 = 1
	version2 = 2
)

func StrTrimSpace(v1str, v2str string) (v1, v2 string) {
	v1 = strings.TrimSpace(v1str)
	v2 = strings.TrimSpace(v2str)
	return
}
func compareSlice(v1slice, v2slice []string) int {
	for index, _ := range v1slice {
		if v1slice[index] > v2slice[index] {
			return version1
		}
		if v1slice[index] < v2slice[index] {
			return version2
		}
		if len(v1slice)-1 == index {
			return version0
		}
	}
	return version0
}

func compareSlice1(v1slice, v2slice []string, flags int) int {
	for index, _ := range v1slice {
		if v1slice[index] > v2slice[index] {
			if flags == 2 {
				return version2
			}
			return version1

		}
		if v1slice[index] < v2slice[index] {
			if flags == 2 {
				return version1
			}
			return version2
		}
		if len(v1slice)-1 == index {
			if flags == 2 {
				return version1
			} else if flags == 1 {
				return version2
			}
		}
	}
	return version0
}

func compareStrVer(v1, v2 string) (res int) {
	s1, s2 := StrTrimSpace(v1, v2)
	v1slice := strings.Split(s1, ".")
	v2slice := strings.Split(s2, ".")
	if len(v1slice) != len(v2slice) {
		if len(v1slice) > len(v2slice) {
			res = compareSlice1(v2slice, v1slice, 2)
			return res
		} else {
			res = compareSlice1(v1slice, v2slice, 1)
			return res
		}
	} else {
		res = compareSlice(v1slice, v2slice)
	}
	return res
}

func demo01() {
	v1 := "5.4.0.1"
	v2 := "5.4.1.1.1.1.1.1.1"
	fmt.Println(compareStrVer(v1, v2))
}
func TestVersions(t *testing.T) {
	demo01()
}
