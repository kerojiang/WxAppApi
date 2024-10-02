/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:ErrorDto.go
 * 修改日期:2024/09/23 10:43:26
 * 作者:kerojiang
 */

package dto

type WXErrorDto struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}
