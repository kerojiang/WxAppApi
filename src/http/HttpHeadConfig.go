/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:HttpHeadConfig.go
 * 修改日期:2024/09/23 16:19:32
 * 作者:kerojiang
 */

package http

type HttpHeadConfig struct {
	Authorization string `json:"Authorization"`
	ContentType   string `json:"Content-Type"`
	Accept        string `json:"Accept"`
}

func NewHttpHeadConfig() *HttpHeadConfig {
	return &HttpHeadConfig{
		Authorization: "",
		ContentType:   "application/json",
		Accept:        "application/json",
	}
}
