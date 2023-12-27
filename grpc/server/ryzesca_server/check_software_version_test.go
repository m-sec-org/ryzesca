package ryzesca_server

import (
	"RyzeSCA/initf"
	"RyzeSCA/internal/model"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestCheckSoftwareVersion(t *testing.T) {
	cline := MysqlCline()
	var cveSoftware []model.CveSoftware
	vendorName := "spring_boot"
	_ = cline.Select([]string{"nssvd_id", "vendor", "product", "version", "cpe", "version_end_excluding",
		"version_end_including", "version_start_excluding", "version_start_including"}).Where("vendor = ?", vendorName).Find(&cveSoftware).Error
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
