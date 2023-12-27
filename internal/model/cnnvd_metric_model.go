package model

type Cnnvd_Metric struct {
	Id           int     `gorm:"primaryKey;column:id"`
	NssvdId      string  `gorm:"column:nssvd_id"`
	CnnvdId      string  `gorm:"column:cnnvd_id"`
	ImpactScore  float64 `gorm:"column:impact_score"`
	AccessVector string  `gorm:"column:access_vevtor"`
}
