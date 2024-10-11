/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:EncryptKey.go
 * 修改日期:2024/10/11 10:20:21
 * 作者:kerojiang
 */

package model

type EncryptKey struct {
	EncryptKey string `json:"encrypt_key"`
	Version    int32  `json:"version"`
	ExpireIn   int32  `json:"expire_in"`
	Iv         string `json:"iv"`
	CreateTime int32  `json:"create_time"`
}
