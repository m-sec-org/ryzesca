package utils

import (
	"fmt"
	"testing"
)

func TestHandleVersionOpenClose(t *testing.T) {
	openClose := HandleVersionOpenClose("2.1.3", "", "", "2.1.0")
	fmt.Print(openClose)
}
