/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:Code2SessionDto.go
 * 修改日期:2024/09/23 10:42:17
 * 作者:kerojiang
 */

package dto

// code2session响应dto
// @Description:
type Code2SessionResDto struct {
	*WXAppErrorDto
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}
