package model

import (
	"RyzeSCA/internal/constant"
	"RyzeSCA/internal/utils"
	"fmt"
	"github.com/package-url/packageurl-go"
	"strings"
)

// Component 在项目内部传递的组件的信息
type Component struct {
	BomRef      string `json:"bom-ref"`
	NameSpace   string `json:"group"`
	PackageName string `json:"name"`
	Purl        string `json:"purl"`
	// 这个是组件传入的type 比如 "type":"library",
	Type string `json:"type"`
	// 组件的类型 比如 maven
	ComponentType string `json:"component_type"`
	Version       string `json:"version"`
}

// GetAllPossibleComponent namespace 要解析进去
func (component *Component) GetAllPossibleComponent() ([]Component, error) {
	// 厂商名称所有可能的组合
	var vendorAliases []string
	// 组件名称所有可能的组合
	var nameAliases []string
	// 返回结果合集
	var result []Component
	//vendor := component.NameSpace
	packageName := component.PackageName
	purl := component.Purl
	pkgType := ""
	var osDistro string
	if purl != "" {
		packageURL, err := packageurl.FromString(component.Purl)
		if err != nil {
			return result, err
		}
		// packageURL 是否是空对象
		pkgType = packageURL.Type
		qualifiers := packageURL.Qualifiers
		distroName := qualifiers.Map()["distro_name"]
		if distroName != "" {
			nameAliases = append(nameAliases, fmt.Sprintf("%s/%s", distroName, packageName))
		}
		osDistro = qualifiers.Map()["distro"]
		if osDistro != "" {
			nameAliases = append(nameAliases, fmt.Sprintf("%s/%s", osDistro, packageName))
		}
		component.NameSpace = packageURL.Namespace
		component.ComponentType = packageURL.Type
		if len(packageName) < 4 {
			var componentA Component
			componentA.NameSpace = component.NameSpace
			componentA.PackageName = component.PackageName
			componentA.Version = component.Version
			componentA.Purl = component.Purl
			componentA.Type = component.Type
			result = append(result, componentA)
			return result, nil
		}
		// 创建python变体
		if strings.HasPrefix(purl, "pkg:pypi") {
			if !strings.HasPrefix(packageName, "python-") {
				nameAliases = append(nameAliases, fmt.Sprintf("python-%s", packageName))
				nameAliases = append(nameAliases, fmt.Sprintf("python-%s_project", packageName))
			}
			vendorAliases = append(vendorAliases, "pip")
			//vendorAliases = append(vendorAliases, "python")
			vendorAliases = append(vendorAliases, fmt.Sprintf("python-%s", packageName))
			vendorAliases = append(vendorAliases, fmt.Sprintf("%sproject", packageName))
		}
		// 创建java变体
		if strings.HasPrefix(purl, "pkg:maven") {
			if strings.Contains(packageName, "-web") {
				nameAliases = append(nameAliases, strings.ReplaceAll(packageName, "-web", "_framework"))
			}
			for _, variant := range constant.CommonSuffixes {
				if strings.Contains(packageName, variant) {
					// 包含 -core 等
					if strings.HasSuffix(packageName, variant) {
						nameAliases = append(nameAliases, strings.ReplaceAll(packageName, variant, ""))
					}
				}
			}
			if strings.Contains(packageName, "-") {
				nameAliases = append(nameAliases, strings.ReplaceAll(packageName, "-", "_"))
			}
		}
		if !utils.SliceFindString(constant.OsPkgTypes, pkgType) {
			for _, suffix := range constant.CommonSuffixes {
				if strings.HasSuffix(packageName, suffix) {
					nameAliases = append(nameAliases, strings.ReplaceAll(packageName, suffix, ""))
				}
			}
		}
		if utils.SliceFindString(constant.OsPkgTypes, pkgType) {
			if strings.Contains(packageName, "lib") {
				nameAliases = append(nameAliases, strings.ReplaceAll(packageName, "lib", ""))
			} else {
				if !strings.Contains(packageName, "lib") {
					nameAliases = append(nameAliases, fmt.Sprintf("%s%s", "lib", packageName))
				}
			}
			if !strings.Contains(packageName, "-bin") {
				nameAliases = append(nameAliases, fmt.Sprintf("%s%s", packageName, "-bin"))
			}
		}
		nameAliases = append(nameAliases, component.PackageName)
		// 去重 组合成Component 返回
		nameAliases = utils.RemoveRepeatedElement(nameAliases)
		vendorAliases = utils.RemoveRepeatedElement(vendorAliases)
		if len(vendorAliases) != 0 {
			for _, vendorAlias := range vendorAliases {
				for _, nameAlias := range nameAliases {
					var componentA Component
					componentA.NameSpace = vendorAlias
					componentA.PackageName = nameAlias
					componentA.Version = component.Version
					componentA.Purl = component.Purl
					componentA.Type = component.Type
					result = append(result, componentA)
				}
			}
		} else {
			for _, nameAlias := range nameAliases {
				var componentA Component
				componentA.NameSpace = component.NameSpace
				componentA.PackageName = nameAlias
				componentA.Version = component.Version
				componentA.Purl = component.Purl
				componentA.Type = component.Type
				result = append(result, componentA)
			}
		}
		return result, nil
	} else {
		return result, nil
	}
}
