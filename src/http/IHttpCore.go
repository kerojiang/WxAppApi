/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:IHttpCore.go
 * 修改日期:2024/09/23 10:05:05
 * 作者:kerojiang
 */

package http

import "wxapp/model/dto"

type IHttpCore interface {
	Get(url string) (*dto.HttpResponseDto, error)

	Post(url string, data string) (*dto.HttpResponseDto, error)
}
