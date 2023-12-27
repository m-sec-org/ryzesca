package model

type CveInfos struct {
	NssvdId          string `gorm:"type:varchar(50);column:nssvd_id" json:"nssvd_id"`
	CveId            string `gorm:"type:varchar(50);column:cve_id" json:"cve_id"`
	CnnvdId          string `gorm:"type:varchar(50);column:cnnvd_id" json:"Cnnvd_id"`
	CnnvdName        string `gorm:"type:varchar(50);column:cnnvd_name" json:"cnnvd_name"`
	CnvdId           string `gorm:"type:varchar(50);column:cnvd_id" json:"cnvd_id"`
	CnvdName         string `gorm:"type:varchar(50);column:cnvd_name" json:"cnvd_name"`
	IsMetric2        int64  `gorm:"type:int;column:is_metric2" json:"is_metric2"`
	IsMetric3        int64  `gorm:"type:int;column:is_metric3" json:"is_metric3"`
	IsCnnvdMetric    int64  `gorm:"type:int;column:is_cnnvd_metric" json:"is_cnnvd_metric"`
	IsCnvdMetric     int64  `gorm:"type:int;column:is_cnvd_metric" json:"is_cnvd_metric"`
	IsRela           int64  `gorm:"type:int;column:is_rela" json:"is_rela"`
	DescriptionEn    string `gorm:"type:text;column:description_en" json:"description_en"`
	DescriptionZh    string `gorm:"type:text;column:description_zh" json:"description_zh"`
	SolutionEn       string `gorm:"type:text;column:solution_en" json:"solution_en"`
	SolutionZh       string `gorm:"type:text;column:solution_zh" json:"solution_zh"`
	PublishedDate    string `gorm:"type:datetime"  json:"published_date"`
	LastModifiedDate string `gorm:"type:datetime"  json:"last_modified_date"`
}
