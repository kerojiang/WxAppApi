/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:IAppBaseApi.go
 * 修改日期:2024/09/23 10:30:50
 * 作者:kerojiang
 */

package wx_api

import "wxapp/model/dto"

// IAppBaseApi 小程序基础接口
// @Description:
type IAppBaseApi interface {
	// 获取接口调用凭证
	//
	// @description:
	//
	// @param:
	//
	// @return:
	GetAccessToken(appId string, secret string) (*dto.AccessTokenResDto, error)

	// Code2Session
	// 小程序登录
	// @description:
	//
	// @param:
	//
	// @return:
	Code2Session(appId string, secret string, code string) (*dto.Code2SessionResDto, error)

	// CheckSessionKey
	// 检查登录状态
	// @description:
	//
	// @param:
	//
	// @return:
	CheckSessionKey(accessToken string, sessionKey string, openid string) (bool, error)

	// ResetUserSessionKey
	// 重置登录状态
	// @description:
	//
	// @param:
	//
	// @return:
	ResetUserSessionKey(accessToken string, sessionKey string,
		openid string) (*dto.Code2SessionResDto, error)
}
