/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:EncryptKeyResDto.go
 * 修改日期:2024/10/10 21:46:43
 * 作者:kerojiang
 */

package dto

import "wxapp/model"

type EncryptKeyResDto struct {
	*WXAppErrorDto
	KeyInfoList []*model.EncryptKey `json:"key_info_list"`
}
