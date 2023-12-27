package utils

import "testing"

func TestSliceFindString(t *testing.T) {
	print(SliceFindString([]string{"a", "b", "c"}, "a"))
}
