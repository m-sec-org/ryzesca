package ryzesca_server

import (
	"RyzeSCA/global"
	"RyzeSCA/grpc/server/ryzesca"
	"RyzeSCA/internal/model"
	"RyzeSCA/internal/utils"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

// ParseNssvdIdPackageNamesToResult 返回最终结果
func ParseNssvdIdPackageNamesToResult(nssvdIdPackageNamesTow [][]model.NssvdIdPackageName) []*ryzesca.PackageVulerability {

	var err error
	var packageVulnerabilities []*ryzesca.PackageVulerability

	//var lock sync.Mutex
	// 返回结果中的 proto 中的  cve_vul_infos

	for _, nssvdIdPackageNames := range nssvdIdPackageNamesTow {
		// 外层循环
		var packageVulnerabilitiesUse ryzesca.PackageVulerability
		var themisCveInfos []*ryzesca.CVEInfo
		for _, nssvdIdPackageName := range nssvdIdPackageNames {

			// 内部拿到每个 NssvdIdPackageName 实例的循环
			// *************************************************************
			// 这里是添加全部的 Package 已经搞定了
			// 下面是 PackageVulnerabilities 需要搞定
			// *************************************************************
			packageVulnerabilitiesUse.PackageName = nssvdIdPackageName.PackageName
			packageVulnerabilitiesUse.PackageVersion = nssvdIdPackageName.Version
			packageVulnerabilitiesUse.Purl = nssvdIdPackageName.PackageUrl
			var packageIdentifiersuse []string
			packageIdentifiersuse = append(packageIdentifiersuse, nssvdIdPackageName.Cpe)
			var vulnerableIdentifiersUse []string
			vulnerableIdentifiersUse = append(vulnerableIdentifiersUse, nssvdIdPackageName.NssvdId)
			var vulnerableVersionsUse []string
			openClose := utils.HandleVersionOpenClose(nssvdIdPackageName.VersionEndExcluding, nssvdIdPackageName.VersionEndIncluding, nssvdIdPackageName.VersionStartExcluding, nssvdIdPackageName.VersionStartIncluding)
			vulnerableVersionsUse = append(vulnerableVersionsUse, openClose)
			// *************************************************************
			// 下面是返回的 cve_vul_infos 里的数据 有很多个 组合到外层循环里
			// *************************************************************
			// 每次内部循环都置空 他在内部循环是唯一值
			var themisCveInfosUse ryzesca.CVEInfo
			// 定义 model.CveInfos
			var modelCveInfosUse model.CveInfos
			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&modelCveInfosUse).Error
			if err != nil {
				fmt.Println("查询失败,结果为空")
				global.Logger.Error("查询失败,结果为空", err)
			}
			// 查询结果转化到 themisCveInfosUse 里面
			themisCveInfosUse.NssvdId = nssvdIdPackageName.NssvdId
			themisCveInfosUse.CveId = modelCveInfosUse.CveId
			themisCveInfosUse.CnnvdId = modelCveInfosUse.CnnvdId
			themisCveInfosUse.CnnvdName = modelCveInfosUse.CnnvdName
			themisCveInfosUse.CnvdId = modelCveInfosUse.CnvdId
			themisCveInfosUse.CnvdName = modelCveInfosUse.CnvdName
			themisCveInfosUse.DescriptionEn = modelCveInfosUse.DescriptionEn
			themisCveInfosUse.DescriptionZh = modelCveInfosUse.DescriptionZh
			themisCveInfosUse.SolutionEn = modelCveInfosUse.SolutionEn
			themisCveInfosUse.SolutionZh = modelCveInfosUse.SolutionZh
			themisCveInfosUse.PublishedDate = modelCveInfosUse.PublishedDate
			themisCveInfosUse.LastModifiedDate = modelCveInfosUse.LastModifiedDate

			var cweIdsUse model.CveRela
			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&cweIdsUse).Error
			if err != nil {
				fmt.Println("查询失败,结果为空")
				global.Logger.Error("查询失败,结果为空", err)
			}
			// 这里是取得地址 等下看一下数据是否正常 不正常的话就要用 初始化一个 themis.CVEInfo_CWE 再去传入地址
			//themisCveInfosUse.CweIds.CweId = cweIdsUse.CweId
			//themisCveInfosUse.CweIds.CweName = cweIdsUse.CweName
			var themisCweIdsUse ryzesca.CVEInfo_CWE
			themisCweIdsUse.CweId = cweIdsUse.CweId
			themisCweIdsUse.CweName = cweIdsUse.CweName
			themisCveInfosUse.CweIds = &themisCweIdsUse
			// identifier 在上面
			themisCveInfosUse.Identifier = nssvdIdPackageName.Cpe
			// version 在上面
			themisCveInfosUse.Version = openClose
			// confidence 为空
			themisCveInfosUse.Identifier = "high"
			var themisCveReferences []*ryzesca.CVEInfo_CVEReference
			var modelCveReferences []model.CveReferences
			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&modelCveReferences).Error
			for _, CveReferences := range modelCveReferences {
				var themisCveReference ryzesca.CVEInfo_CVEReference
				themisCveReference.Url = CveReferences.Url
				themisCveReference.Name = CveReferences.Name
				themisCveReference.Refsource = CveReferences.Reference
				themisCveReference.Tags = CveReferences.Tags
				themisCveReferences = append(themisCveReferences, &themisCveReference)
			}
			themisCveInfosUse.CveReferences = themisCveReferences

			var cveMetric2use model.CveMetric2
			// 创建一个model 里的 CveMetric2
			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&cveMetric2use).Error
			//
			var cveMetric2Map map[string]*structpb.Value
			cveMetric2Map = make(map[string]*structpb.Value)
			valMetrica := reflect.ValueOf(cveMetric2use) //获取reflect.Type类型
			typMetric2 := reflect.TypeOf(cveMetric2use)
			for i := 0; i < valMetrica.NumField(); i++ {
				//if
				if valMetrica.Field(i).Kind() == reflect.Float64 {
					if valMetrica.Field(i).Float() == 0 {
						cveMetric2Map[typMetric2.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cveMetric2Map[typMetric2.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								NumberValue: valMetrica.Field(i).Float(),
							},
						}
					}
				}
				if valMetrica.Field(i).Kind() == reflect.String {
					if valMetrica.Field(i).String() == "" {
						cveMetric2Map[typMetric2.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cveMetric2Map[typMetric2.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_StringValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								StringValue: valMetrica.Field(i).String(),
							},
						}
					}
				}
			}
			themisCveInfosUse.CveMetric2 = cveMetric2Map

			var cveMetric3use model.CveMetric3

			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&cveMetric3use).Error
			var cveMetric3Map map[string]*structpb.Value
			cveMetric3Map = make(map[string]*structpb.Value)
			valMetric3 := reflect.ValueOf(cveMetric3use) //获取reflect.Type类型
			typMetric3 := reflect.TypeOf(cveMetric3use)
			for i := 0; i < valMetric3.NumField(); i++ {
				if valMetric3.Field(i).Kind() == reflect.Float64 {
					if valMetric3.Field(i).Float() == 0 {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								NumberValue: valMetric3.Field(i).Float(),
							},
						}
					}
				}
				if valMetric3.Field(i).Kind() == reflect.String {
					if valMetric3.Field(i).String() == "" {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_StringValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								StringValue: valMetrica.Field(i).String(),
							},
						}
					}
				}
				if valMetric3.Field(i).Kind() == reflect.Int64 {
					// int64 先转成 float64 再转成 string
					if valMetric3.Field(i).Int() == 0 {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cveMetric3Map[typMetric3.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: float64(valMetrica.Field(i).Int()),
							},
						}
					}
				}
			}
			themisCveInfosUse.CveMetric3 = cveMetric3Map

			var cnnvdMetricUse model.Cnnvd_Metric
			err = global.MysqlDB.Select("*").Where("nssvd_id = ?", nssvdIdPackageName.NssvdId).Find(&cnnvdMetricUse).Error
			var cnnvdMetricMap map[string]*structpb.Value
			cnnvdMetricMap = make(map[string]*structpb.Value)
			valMetricCnnvd := reflect.ValueOf(cnnvdMetricUse) //获取reflect.Type类型
			typMetricCnnvd := reflect.TypeOf(cnnvdMetricUse)
			for i := 0; i < valMetricCnnvd.NumField(); i++ {
				if valMetricCnnvd.Field(i).Kind() == reflect.Float64 {
					if valMetricCnnvd.Field(i).Float() == 0 {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								NumberValue: valMetricCnnvd.Field(i).Float(),
							},
						}
					}
				}
				if valMetricCnnvd.Field(i).Kind() == reflect.String {
					if valMetricCnnvd.Field(i).String() == "" {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_StringValue{
								// 判断一下 valMetrica.Field(i).Interface() 的类型
								StringValue: valMetrica.Field(i).String(),
							},
						}
					}
				}
				if valMetricCnnvd.Field(i).Kind() == reflect.Int64 {
					// int64 先转成 float64 再转成 string
					if valMetricCnnvd.Field(i).Int() == 0 {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_NullValue{
								NullValue: structpb.NullValue_NULL_VALUE,
							},
						}
					} else {
						cnnvdMetricMap[typMetricCnnvd.Field(i).Tag.Get("json")] = &structpb.Value{
							Kind: &structpb.Value_StringValue{
								StringValue: fmt.Sprintf("%f", float64(valMetricCnnvd.Field(i).Int())),
							},
						}
					}
				}
			}
			themisCveInfosUse.CnnvdMetric = cnnvdMetricMap

			// risk_level
			level := utils.GetRiskLevel(cveMetric2use.BaseScore, cveMetric3use.BaseScore, cnnvdMetricUse.ImpactScore)
			themisCveInfosUse.RiskLevel = int32(level)
			themisCveInfos = append(themisCveInfos, &themisCveInfosUse)
			packageVulnerabilitiesUse.CveVulInfos = themisCveInfos
			packageVulnerabilitiesUse.PackageIdentifiers = packageIdentifiersuse
			packageVulnerabilitiesUse.VulnerableIdentifiers = vulnerableIdentifiersUse
			packageVulnerabilitiesUse.VulnerableVersions = vulnerableVersionsUse
			// 里层循环
		}
		packageVulnerabilities = append(packageVulnerabilities, &packageVulnerabilitiesUse)
	}

	return packageVulnerabilities
}
