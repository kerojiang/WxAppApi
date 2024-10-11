/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:PhoneInfoResDto.go
 * 修改日期:2024/10/10 22:08:02
 * 作者:kerojiang
 */

package dto

import "wxapp/model"

type PhoneInfoResDto struct {
	*WXAppErrorDto
	Phoneinfo *model.PhoneInfo `json:"phone_info"`
}
