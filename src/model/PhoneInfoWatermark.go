/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:PhoneInfoWatermark.go
 * 修改日期:2024/10/11 10:19:58
 * 作者:kerojiang
 */

package model

type PhoneInfoWatermark struct {
	Timestamp int32  `json:"timestamp"`
	AppId     string `json:"appid"`
}
