package utils

import "RyzeSCA/internal/constant"

// GetRiskLevel risk 的等级问题
func GetRiskLevel(cvss2Score float64, cvss3Score float64, cnnvdScore float64) constant.VulRiskLevel {
	if cvss3Score != 0 {
		if cvss3Score >= 0 && cvss3Score <= 4 {
			// 低危
			return constant.LOW
		} else {
			if cvss3Score >= 4.1 && cvss3Score <= 6.9 {
				return constant.MEDIUM
			}
			if cvss3Score >= 7.0 && cvss3Score <= 8.9 {
				return constant.HIGH
			}
			if cvss3Score >= 9.0 && cvss3Score <= 10.0 {
				return constant.CRITICAL
			}
		}
		return constant.NOTHING
	} else {
		if cvss2Score != 0 {
			if cvss2Score >= 0 && cvss2Score <= 4 {
				// 低危
				return constant.LOW
			} else {
				if cvss2Score >= 4.1 && cvss2Score <= 6.9 {
					return constant.MEDIUM
				}
				if cvss2Score >= 7.0 && cvss2Score <= 8.9 {
					return constant.HIGH
				}
			}
			return constant.NOTHING
		} else {
			if cnnvdScore != 0 {
				if cnnvdScore > 0 && cnnvdScore < 2 {
					return constant.LOW
				}
				if cnnvdScore >= 2 && cnnvdScore < 4 {
					return constant.MEDIUM
				}
				if cnnvdScore >= 4 && cnnvdScore < 6 {
					return constant.HIGH
				}
				return constant.NOTHING
			} else {
				return constant.MEDIUM
			}
		}
	}
}
