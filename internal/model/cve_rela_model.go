package model

type CveRela struct {
	NssvdId string `gorm:"column:nssvd_id" json:"nssvd_Id"`
	CweId   string `gorm:"column:cwe_id" json:"cwe_id"`
	CweName string `gorm:"column:cwe_name" json:"cwe_name"`
}
