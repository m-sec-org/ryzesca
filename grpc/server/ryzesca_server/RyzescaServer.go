package ryzesca_server

import (
	"RyzeSCA/global"
	"RyzeSCA/grpc/server/ryzesca"
	"RyzeSCA/internal/model"
	"RyzeSCA/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
	"time"
)

var errUser error

type RequestCyclonedxJsonpath struct {
	Schema     string            `json:"$schema"`
	BomFormat  string            `json:"bomFormat"`
	Components []model.Component `json:"components"`
}
type RyzescaServer struct {
}

func (RyzescaServer RyzescaServer) RunRyzescaCycloneDX(ctx context.Context, request *ryzesca.RyzescaParams) (*ryzesca.RyzescaResult, error) {
	startTime := time.Now()
	startTimestamp := startTime.Format("2006-01-02 15:04:05")
	RequestCyclonedxJsonpath := RequestCyclonedxJsonpath{}
	errUser = json.Unmarshal([]byte(request.CyclonedxJson), &RequestCyclonedxJsonpath)
	if errUser != nil {
		fmt.Println("转json失败")
		endTime := time.Now()
		endTimestamp := endTime.Format("2006-01-02 15:04:05")
		return &ryzesca.RyzescaResult{
			Headers: &ryzesca.Header{
				ToolName:       "themis",
				ToolVersion:    "",
				StartTimestamp: startTimestamp,
				EndTimestamp:   endTimestamp,
				Duration:       startTime.Sub(endTime).Seconds(),
				Message:        "json解析失败",
			},
			Files: nil,
		}, nil
	}
	var nssvdIdPackageNamesTwo [][]model.NssvdIdPackageName
	var packages []*ryzesca.Package
	for _, component := range RequestCyclonedxJsonpath.Components {
		packagesUse := ryzesca.Package{
			Name:      component.PackageName,
			Type:      component.Type,
			Version:   component.Version,
			Purl:      component.Purl,
			Namespace: component.NameSpace,
			DeclaredLicense: &structpb.Value{
				Kind: &structpb.Value_NullValue{
					NullValue: structpb.NullValue_NULL_VALUE,
				},
			},
		}
		packages = append(packages, &packagesUse)
		var componentSelect []model.CveSoftware
		allPossibleComponents, _ := component.GetAllPossibleComponent()
		if len(allPossibleComponents) == 0 {
			continue
		} else {
			for _, possibleComponent := range allPossibleComponents {
				// 每个可能的组件的搜索结果 结果是个列表
				selectComponent := SelectComponent(possibleComponent)
				if len(selectComponent) == 0 {
					// 查询结果为空
					continue
				} else {
					empty := model.RemoveRepeatedElementAndEmpty(selectComponent)
					//fmt.Println(selectComponent)
					version := CheckSoftwareVersion(empty, possibleComponent.Version)
					if len(version) != 0 {
						//fmt.Println(possibleComponent)
						//fmt.Println(version)
						//fmt.Println(len(version))
						componentSelect = append(componentSelect, version...)
					}
					//fmt.Println(json.Marshal(selectComponent))
					//marshal, _ := json.Marshal(selectComponent)
					//fmt.Println(string(marshal))
				}
			}

		}
		componentSelect = model.RemoveRepeatedElement(componentSelect)
		if len(componentSelect) != 0 {
			//fmt.Println()
			//fmt.Println("***********************")
			//fmt.Println(component)
			//fmt.Println(componentSelect)
			//fmt.Println(len(componentSelect))
			//fmt.Println("***********************")
			//fmt.Println()
			var nssvdIdPackageNames []model.NssvdIdPackageName
			for _, software := range componentSelect {
				nssvdIdPackageNames = append(nssvdIdPackageNames, model.NssvdIdPackageName{
					NssvdId:               software.NssvdId,
					PackageName:           software.Product,
					Version:               component.Version,
					PackageUrl:            component.Purl,
					Type:                  component.Type,
					NameSpace:             component.NameSpace,
					Cpe:                   software.Cpe,
					VersionEndExcluding:   software.VersionEndExcluding,
					VersionEndIncluding:   software.VersionEndIncluding,
					VersionStartExcluding: software.VersionStartExcluding,
					VersionStartIncluding: software.VersionStartIncluding,
				})
			}
			nssvdIdPackageNamesTwo = append(nssvdIdPackageNamesTwo, nssvdIdPackageNames)
		}
	}
	PackageVulerability := ParseNssvdIdPackageNamesToResult(nssvdIdPackageNamesTwo)
	// 初始化file数据信息
	file := ryzesca.File{
		Path:     "sbom",
		Name:     "sbom",
		Type:     "file",
		BaseName: "sbom",
	}
	file.Packages = packages
	file.PackageVulnerabilities = PackageVulerability
	endTime := time.Now()
	endTimestamp := endTime.Format("2006-01-02 15:04:05")
	return &ryzesca.RyzescaResult{
		Headers: &ryzesca.Header{
			ToolName:       "themis",
			ToolVersion:    "",
			StartTimestamp: startTimestamp,
			EndTimestamp:   endTimestamp,
			Duration:       startTime.Sub(endTime).Seconds(),
			Message:        "json解析失败",
		},
		Files: &file,
	}, nil
}

func SelectComponent(component model.Component) []model.CveSoftware {
	var vendorName string
	var cveSoftware []model.CveSoftware
	if component.PackageName != "" {
		if component.NameSpace != "" {
			NameSpaceList := strings.Split(component.NameSpace, ".")
			if len(NameSpaceList) < 1 {
				vendorName = NameSpaceList[1]
			} else {
				vendorName = NameSpaceList[0]
			}
		} else {
			vendorName = ""
		}
	}
	if vendorName != "" {
		_ = global.MysqlDB.Select([]string{"nssvd_id", "vendor", "product", "version", "cpe", "version_end_excluding",
			"version_end_including", "version_start_excluding", "version_start_including"}).Where("vendor = ?", vendorName).Find(&cveSoftware).Error
		if len(cveSoftware) != 0 {
			return cveSoftware
		} else {
			errUser = global.MysqlDB.Select([]string{"nssvd_id", "cpe", "version_end_excluding",
				"version_end_including", "version_start_excluding", "version_start_including"}).Where("Product = ? and vulnerable = ?", component.PackageName, 1).Find(&cveSoftware).Error
			return cveSoftware
		}
	} else {
		errUser = global.MysqlDB.Select([]string{"nssvd_id", "cpe", "version_end_excluding",
			"version_end_including", "version_start_excluding", "version_start_including"}).Where("Product = ? and vulnerable = ?", component.PackageName, 1).Find(&cveSoftware).Error
		return cveSoftware
	}

}

func CheckSoftwareVersion(software []model.CveSoftware, targetVersion string) []model.CveSoftware {
	// 返回结果
	var result []model.CveSoftware
	for _, cveSoftware := range software {
		if cveSoftware.VersionStartIncluding != "" && cveSoftware.VersionEndIncluding != "" {
			// 开始和结束都不为空时候 只要版本相等就符合条件
			if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 0 || utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 0 {
				result = append(result, cveSoftware)
			} else {
				if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 2 {
					result = append(result, cveSoftware)
				}
			}
		} else {
			if cveSoftware.VersionStartExcluding != "" && cveSoftware.VersionEndExcluding != "" {
				// 开始和结束都不为空时候 只要版本相等就符合条件
				if utils.VersionComparison(targetVersion, cveSoftware.VersionStartExcluding) == 1 || utils.VersionComparison(targetVersion, cveSoftware.VersionEndExcluding) == 2 {
					result = append(result, cveSoftware)
				}
			} else {
				if cveSoftware.VersionStartExcluding != "" && cveSoftware.VersionEndIncluding != "" {
					// 大于开始版本号 小于等于结束版本号
					if utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 0 {
						result = append(result, cveSoftware)
					} else {
						if utils.VersionComparison(targetVersion, cveSoftware.VersionStartExcluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 2 {
							result = append(result, cveSoftware)
						}
					}
				} else {
					if cveSoftware.VersionStartIncluding != "" && cveSoftware.VersionEndExcluding != "" {
						// 等于开始包含的版本号
						if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 0 {
							result = append(result, cveSoftware)
						} else {
							if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndExcluding) == 2 {
								result = append(result, cveSoftware)
							}
						}
					} else {

						if cveSoftware.VersionStartExcluding != "" {
							if utils.VersionComparison(targetVersion, cveSoftware.VersionStartExcluding) == 1 {
								result = append(result, cveSoftware)
							}
						}
						if cveSoftware.VersionStartIncluding != "" {
							if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 1 || utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 0 {
								result = append(result, cveSoftware)
							}
						}
						if cveSoftware.VersionEndExcluding != "" {
							if utils.VersionComparison(targetVersion, cveSoftware.VersionEndExcluding) == 2 {
								result = append(result, cveSoftware)
							}
						}
						if cveSoftware.VersionEndIncluding != "" {
							if utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 2 || utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 0 {
								result = append(result, cveSoftware)
							}
						}
					}
				}
			}
		}

	}
	return result
}
