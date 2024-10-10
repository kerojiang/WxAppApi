/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:PhoneInfoResDto.go
 * 修改日期:2024/10/10 22:08:02
 * 作者:kerojiang
 */

package dto

type PhoneInfoResDto struct {
	WXAppErrorDto
	Phoneinfo PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string             `json:"phoneNumber"`
	PurePhoneNumber string             `json:"purePhoneNumber"`
	CountryCode     string             `json:"countryCode"`
	Watermark       PhoneInfoWatermark `json:"watermark"`
}

type PhoneInfoWatermark struct {
	Timestamp int32  `json:"timestamp"`
	AppId     string `json:"appid"`
}
