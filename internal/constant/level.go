package constant

// VulRiskLevel Level 等级常量
type VulRiskLevel int32

const (
	NOTHING  VulRiskLevel = 0
	LOW      VulRiskLevel = 1
	MEDIUM   VulRiskLevel = 2
	HIGH     VulRiskLevel = 3
	CRITICAL VulRiskLevel = 4
)

// Confidence 组件风险等级常量
type Confidence string

const (
	ConfigUNKNOWN Confidence = "unknown"
	ConfigLOW     Confidence = "low"
	ConfigMedium  Confidence = "medium"
	ConfigHIGH    Confidence = "high"
	ConfigHIGHEST Confidence = "highest"
)
