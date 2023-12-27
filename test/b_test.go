package test

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"testing"
)

func TestB(t *testing.T) {
	instance, _ := packageurl.FromString("pkg:maven/javax.xml.bind/jaxb-api@2.3.1?type=jar")
	fmt.Println(instance)

	fmt.Println(instance.Namespace)
	fmt.Println(instance.Type)
}
