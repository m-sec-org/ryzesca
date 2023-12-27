package model

// NssvdIdPackageName 用来传递nssvd_id 和 package_name version purl 的结构体
type NssvdIdPackageName struct {
	NssvdId               string `json:"nssvd_id"`
	PackageName           string `json:"package"`
	Version               string `json:"version"`
	PackageUrl            string `json:"package_url"`
	Type                  string `json:"type"`
	NameSpace             string `json:"name_space"`
	Cpe                   string `json:"cpe"`
	VersionEndExcluding   string `json:"version_end_excluding"`
	VersionEndIncluding   string `json:"version_end_including"`
	VersionStartExcluding string `json:"version_start_excluding"`
	VersionStartIncluding string `json:"version_start_including"`
}
