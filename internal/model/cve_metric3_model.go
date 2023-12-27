package model

type CveMetric3 struct {
	ID                    int64   `gorm:"primary_key;column:id" json:"id"`
	NssvdId               string  `gorm:"column:nssvd_id" json:"nssvd_Id"`
	ExploitabilityScore   float64 `gorm:"column:exploitability_score" json:"exploitability_score"`
	ImpactScore           float64 `gorm:"column:impact_score" json:"impact_score"`
	Version               string  `gorm:"column:version" json:"version"`
	VectorString          string  `gorm:"column:vector_string" json:"vector_string"`
	AttackVector          string  `gorm:"column:attack_vector" json:"attack_vector"`
	AttackComplexity      string  `gorm:"column:attack_complexity" json:"attack_complexity"`
	PrivilegesRequired    string  `gorm:"column:privileges_required" json:"privileges_required"`
	UserInteraction       string  `gorm:"column:user_interaction" json:"user_interaction"`
	Scope                 string  `gorm:"column:scope" json:"scope"`
	ConfidentialityImpact string  `gorm:"column:confidentiality_impact" json:"confidentiality_impact"`
	IntegrityImpact       string  `gorm:"column:integrity_impact" json:"integrity_impact"`
	AvailabilityImpact    string  `gorm:"column:availability_impact" json:"availability_impact"`
	BaseScore             float64 `gorm:"column:base_score" json:"base_score"`
	BaseSeverity          string  `gorm:"column:base_severity" json:"base_severity"`
}
