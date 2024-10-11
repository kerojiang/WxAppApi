/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:PhoneInfo.go
 * 修改日期:2024/10/11 10:19:39
 * 作者:kerojiang
 */

package model

type PhoneInfo struct {
	PhoneNumber     string              `json:"phoneNumber"`
	PurePhoneNumber string              `json:"purePhoneNumber"`
	CountryCode     string              `json:"countryCode"`
	Watermark       *PhoneInfoWatermark `json:"watermark"`
}
