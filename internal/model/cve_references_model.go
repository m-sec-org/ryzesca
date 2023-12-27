package model

type CveReferences struct {
	ID        int64  `gorm:"primary_key;column:id" json:"id"`
	NssvdId   string `gorm:"column:nssvd_id" json:"nssvd_Id"`
	Reference string `gorm:"column:reference" json:"reference"`
	Name      string `gorm:"column:name" json:"name"`
	Url       string `gorm:"column:url" json:"url"`
	Tags      string `gorm:"column:tags" json:"tags"`
}
