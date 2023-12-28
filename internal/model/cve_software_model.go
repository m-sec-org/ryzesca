package model

type CveSoftware struct {
	ID                    int64  `gorm:"primary_key;column:id" json:"id"`
	NssvdId               string `gorm:"column:nssvd_id" json:"nssvd_Id"`
	Vendor                string `gorm:"column:vendor" json:"vendor"`
	Product               string `gorm:"column:product" json:"product"`
	Version               string `gorm:"column:version" json:"version"`
	UpdateVersion         string `gorm:"column:update_version" json:"update_version"`
	Cpe                   string `gorm:"column:cpe" json:"cpe"`
	VersionEndExcluding   string `gorm:"column:version_end_excluding" json:"version_end_excluding"`
	VersionEndIncluding   string `gorm:"column:version_end_including" json:"version_end_including"`
	VersionStartExcluding string `gorm:"column:version_start_excluding" json:"version_start_excluding"`
	VersionStartIncluding string `gorm:"column:version_start_including" json:"version_start_including"`
	Vulnerable            int    `gorm:"column:vulnerable" json:"vulnerable"`
}

// RemoveRepeatedElementAndEmpty ,去除重复的元素和空的元素
func RemoveRepeatedElementAndEmpty(slice []CveSoftware) []CveSoftware {
	// 去除重复的元素
	//slice = RemoveRepeatedElement(slice)
	// 去除空的元素
	slice = RemoveEmpty(slice)
	return slice
}
func RemoveRepeatedElement(slice []CveSoftware) []CveSoftware {
	newSlice := make([]CveSoftware, 0)
	for i := 0; i < len(slice); i++ {
		repeat := false
		for j := i + 1; j < len(slice); j++ {
			if slice[i].NssvdId == slice[j].NssvdId {
				repeat = true
				break
			}
		}
		if !repeat {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func RemoveEmpty(slice []CveSoftware) []CveSoftware {
	newSlice := make([]CveSoftware, 0)
	for i := 0; i < len(slice); i++ {
		if slice[i].NssvdId != "" {
			if slice[i].VersionStartExcluding != "" || slice[i].VersionStartIncluding != "" || slice[i].VersionEndExcluding != "" || slice[i].VersionEndIncluding != "" {
				newSlice = append(newSlice, slice[i])
			}
		}
	}
	return newSlice
}
