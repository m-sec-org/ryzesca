package test

import (
	"RyzeSCA/initf"
	"RyzeSCA/internal/model"
	"RyzeSCA/internal/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestC(t *testing.T) {
	cline := MysqlCline()
	var cveSoftware []model.CveSoftware
	vendorName := "spring_boot"
	_ = cline.Select([]string{"nssvd_id", "vendor", "product", "version", "cpe", "version_end_excluding",
		"version_end_including", "version_start_excluding", "version_start_including"}).Where("Product = ?", vendorName).Find(&cveSoftware).Error
	version := CheckSoftwareVersion(cveSoftware, "2.2.2.RELEASE")
	fmt.Println(version)
}
func MysqlCline() *gorm.DB {
	initf.ViperInit()
	mysql, err := initf.MysqlInit()
	if err != nil {
		return nil
	}
	return mysql
}
func CheckSoftwareVersion(software []model.CveSoftware, targetVersion string) []model.CveSoftware {
	// 返回结果
	var result []model.CveSoftware
	for _, cveSoftware := range software {
		marshal, _ := json.Marshal(cveSoftware)
		fmt.Println(string(marshal))
		if cveSoftware.VersionStartIncluding != "" && cveSoftware.VersionEndIncluding != "" {
			// 开始和结束都不为空时候 只要版本相等就符合条件
			if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 0 || utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 1 {
				result = append(result, cveSoftware)
			} else {
				if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 2 {
					result = append(result, cveSoftware)
				}
			}
		}
		if cveSoftware.VersionStartExcluding != "" && cveSoftware.VersionEndExcluding != "" {
			// 开始和结束都不为空时候 只要版本相等就符合条件
			if utils.VersionComparison(targetVersion, cveSoftware.VersionStartExcluding) == 1 || utils.VersionComparison(targetVersion, cveSoftware.VersionEndExcluding) == 2 {
				result = append(result, cveSoftware)
			}
		}
		if cveSoftware.VersionStartExcluding != "" && cveSoftware.VersionEndIncluding != "" {
			// 大于开始版本号 小于等于结束版本号
			if utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 0 {
				result = append(result, cveSoftware)
			} else {
				if utils.VersionComparison(targetVersion, cveSoftware.VersionStartExcluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndIncluding) == 2 {
					result = append(result, cveSoftware)
				}
			}
		}
		if cveSoftware.VersionStartIncluding != "" && cveSoftware.VersionEndExcluding != "" {
			// 等于开始包含的版本号
			if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 0 {
				result = append(result, cveSoftware)
			} else {
				if utils.VersionComparison(targetVersion, cveSoftware.VersionStartIncluding) == 1 && utils.VersionComparison(targetVersion, cveSoftware.VersionEndExcluding) == 2 {
					result = append(result, cveSoftware)
				}
			}
		}
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
				fmt.Println("suss")
				result = append(result, cveSoftware)
			}
		}

	}
	return result
}
