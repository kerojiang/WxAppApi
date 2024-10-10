/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:EncryptKeyResDto.go
 * 修改日期:2024/10/10 21:46:43
 * 作者:kerojiang
 */

package dto

type EncryptKeyResDto struct {
	WXAppErrorDto
	KeyInfoList *[]EncryptKetDto `json:"key_info_list"`
}

type EncryptKetDto struct {
	EncryptKey string `json:"encrypt_key"`
	version    int32  `json:"version"`
	ExpireIn   int32  `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int32  `json:"create_time"`
}
