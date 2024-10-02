/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:AccessTokenDto.go
 * 修改日期:2024/09/23 10:32:50
 * 作者:kerojiang
 */

package dto

type AccessTokenResDto struct {
	WXAppErrorDto
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
