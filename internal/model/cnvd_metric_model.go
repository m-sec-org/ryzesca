package model

type Cnvd_Metric struct {
	Id                int     `gorm:"primaryKey;column:id"`
	Nssvd_id          string  `gorm:"column:nssvd_id"`
	Cnvd_id           string  `gorm:"column:cnvd_id"`
	Impact_score      float64 `gorm:"column:impact_score"`
	Access_vector     string  `gorm:"column:access_vevtor"`
	Patch_name        string  `gorm:"column:patch_name"`
	Patch_description string  `gorm:"column:patch_description"`
}
