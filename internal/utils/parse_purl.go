package utils

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"regexp"
	"strings"
)

// ParsePurl 针对purl进行解析
func ParsePurl(purl string) (PurlObj, error) {
	var purlObj PurlObj
	purlObj.Purl = purl
	// 判断是否为空 必须包含@ 必须包含/
	if len(strings.TrimSpace(purl)) == 0 || !strings.Contains(purl, "@") || !strings.Contains(purl, "/") {
		return purlObj, fmt.Errorf("purl is empty")
	}

	instance, err := packageurl.FromString(purl)
	if err == nil {
		// packageurl-go包解析成功
		return PurlObj{
			Purl:      purl,
			Namespace: instance.Namespace,
			Name:      instance.Name,
			Version:   instance.Version,
		}, nil
	} else {
		// packageurl-go包解析失败 走自己的解析逻辑
		version, err := ParseVersion(purl)
		if err != nil {
			version = ""
		}
		// 判断是否存在@
		if !strings.Contains(purl, "@") {
			return purlObj, fmt.Errorf("purl is not contains @")
		}
		// "purl":"pkg:pypi/pycrypto@2.6.1" 解析出来版本号
		tmpA := strings.Split(purl, "@")[0]
		tmpB := strings.Split(tmpA, "/")
		// java 跟 golang的组件属于一种
		// python 跟其他的属于一种
		if len(tmpB) == 2 {
			purlObj.Namespace = tmpB[0]
			purlObj.Name = tmpB[1]
			purlObj.Version = version
		}
		if len(tmpB) > 2 {
			purlObj.Name = tmpB[len(tmpB)-1]
			namespaceA := tmpB[len(tmpB)-2]
			// 判断是否以pkg:开头
			if strings.HasPrefix(namespaceA, "pkg:") {
				namespaceB := strings.Split(namespaceA, ":")
				namespaceC := namespaceB[len(namespaceB)-1]
				purlObj.Namespace = namespaceC
			} else {
				purlObj.Namespace = tmpB[len(tmpB)-2]
			}
			purlObj.Name = tmpB[len(tmpB)-1]
		}
		return purlObj, nil
	}
}

func ParseVersion(purl string) (string, error) {
	if len(strings.TrimSpace(purl)) == 0 || !strings.Contains(purl, "@") || !strings.Contains(purl, "/") {
		return "", fmt.Errorf("purl is wrong")
	}
	// 判断是否存在?
	if strings.Contains(purl, "?") {
		// 存在? 说明版本号在 @ 和 ? 之间
		// 使用正则提取 @ 和 ? 之间的内容
		compile := regexp.MustCompile(`@(.*?)\?`)
		version := compile.FindString(purl)
		version = strings.Replace(version, "@", "", -1)
		version = strings.Replace(version, "?", "", -1)
		return version, nil
	} else {
		// 不存在? 说明版本号在 @ 之后
		// 使用正则提取 @ 之后的内容
		version := strings.Split(purl, "@")[1]
		return version, nil
	}
}
