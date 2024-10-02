/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:PayUnionidResDto.go
 * 修改日期:2024/10/02 23:26:31
 * 作者:kerojiang
 */

package dto

type PayUnionidResDto struct {
	WXAppErrorDto
	UnionId string `json:"unionid"`
}
