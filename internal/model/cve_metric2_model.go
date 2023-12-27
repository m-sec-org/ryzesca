package model

type CveMetric2 struct {
	ID                      int64   `gorm:"primary_key;column:id" json:"id"`
	NssvdId                 string  `gorm:"column:nssvd_id" json:"nssvd_Id"`
	Severity                string  `gorm:"column:severity" json:"severity"`
	ExploitabilityScore     float64 `gorm:"column:exploitability_score" json:"exploitability_score"`
	ImpactScore             float64 `gorm:"column:impact_score" json:"impact_score"`
	AcInsufInfo             float64 `gorm:"column:ac_insuf_info" json:"ac_insuf_info"`
	ObtainAllPrivilege      float64 `gorm:"column:obtain_all_privilege" json:"obtain_all_privilege"`
	ObtainUserPrivilege     float64 `gorm:"column:obtain_user_privilege" json:"obtain_user_privilege"`
	ObtainOtherPrivilege    float64 `gorm:"column:obtain_other_privilege" json:"obtain_other_privilege"`
	UserInteractionRequired float64 `gorm:"column:user_interaction_required" json:"user_interaction_required"`
	Version                 string  `gorm:"column:version" json:"version"`
	VectorString            string  `gorm:"column:vector_string" json:"vector_string"`
	AccessVector            string  `gorm:"column:access_vector" json:"access_vector"`
	AccessComplexity        string  `gorm:"column:access_complexity" json:"access_complexity"`
	Authentication          string  `gorm:"column:authentication" json:"authentication"`
	ConfidentialityImpact   string  `gorm:"column:confidentiality_impact" json:"confidentiality_impact"`
	IntegrityImpact         string  `gorm:"column:integrity_impact" json:"integrity_impact"`
	AvailabilityImpact      string  `gorm:"column:availability_impact" json:"availability_impact"`
	BaseScore               float64 `gorm:"column:base_score" json:"base_score"`
}
