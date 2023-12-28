package utils

import (
	"fmt"
	"testing"
)

func demo01() {
	v1 := "5.4.0.4"
	v2 := "5.4.0.5"

	v3 := "1.11.3"
	v4 := "1.8.3"
	v5 := "1.12.3.4"
	v6 := "1.12.3.4.0"
	fmt.Println(VersionComparison(v1, v2))
	fmt.Println(VersionComparison(v3, v4))
	fmt.Println(VersionComparison(v5, v6))
}
func TestVersionComparison(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5"}
	s2 := []string{"6", "7", "8"}
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	le := len(s1) - len(s2)
	for i := 0; i < le; i++ {
		fmt.Println("*")
		s2 = append(s2, "0")
	}
	fmt.Println(len(s1))
	fmt.Println(len(s2))
}
