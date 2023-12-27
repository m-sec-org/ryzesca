package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	name := "spring-boot-starter-json"
	vendor := "org.springframework.boot"
	tmpA := strings.Split(vendor, ".")
	if len(tmpA) > 1 {
		if len(tmpA) > 3 {
			if tmpA[1] != name {
				fmt.Println(tmpA[1])
			}
		}
	}
}
