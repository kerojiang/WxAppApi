/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:IAppBaseApi.go
 * 修改日期:2024/09/23 10:30:50
 * 作者:kerojiang
 */

package wx_api

import "wxapp/model/dto"

type IAppBaseApi interface {
	//获取接口调用凭证
	//
	// @description:
	//
	// @param:
	//
	// @return:
	GetAccessToken(appId string, secret string) (dto.AccessTokenDto, error)

	// 小程序登录
	//
	// @description:
	//
	// @param:
	//
	// @return:
	Code2Session(appId string, secret string, code string) (dto.Code2SessionDto, error)

	// 检查登录状态
	//
	// @description:
	//
	// @param:
	//
	// @return:
	CheckSessionKey(accessToken string, openid string) (bool, error)

	// 重置登录状态
	//
	// @description:
	//
	// @param:
	//
	// @return:
	ResetUserSessionKey(accessToken string, openid string) (dto.Code2SessionDto, error)
}
