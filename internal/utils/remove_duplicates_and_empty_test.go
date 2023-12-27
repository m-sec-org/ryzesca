package utils

import (
	"fmt"
	"testing"
)

func TestRemoveRepeatedElement(t *testing.T) {
	ret := RemoveRepeatedElement([]string{"a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c"})
	fmt.Println(ret)
}
