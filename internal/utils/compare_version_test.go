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
	//demo01()
	//s1 := []string{"1", "2", "3", "4", "5"}
	//s2 := []string{"6", "7", "8"}
	//for i := 0; i < len(s1)-len(s2)+1; i++ {
	//	s2 = append(s2, "0")
	//}
	//fmt.Println(len(s1))
	//fmt.Println(len(s2))
	//fmt.Println(len(s1) - len(s2))
	//fmt.Println(s2)
	//fmt.Println(s1)
	//v1, _ := strconv.ParseUint(strconv.Itoa(0), 10, 64)
	//atoi, _ := strconv.Atoi("0")
	//fmt.Println(v1)
	//fmt.Println(atoi)
	//v1str := "1.23635.56.5156.65.515.626fafasdf.155.515asdf.51afsd.asasdf"
	//res, _ := regexp.Compile(constant.VERSIONRE)
	//matchString := res.FindStringSubmatch(strings.TrimSpace(v1str))
	////match := res.Match([]byte(strings.TrimSpace(v1str)))
	//fmt.Println(matchString)

	//comparison := VersionComparison("1.2.3.4.56.5", "13.25.1")
	//comparison1 := VersionComparison("13.25.1", "1.2.3.4.56.5")
	//fmt.Println(comparison)
	//fmt.Println(comparison1)

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
