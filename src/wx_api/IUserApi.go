/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:IUserApi.go
 * 修改日期:2024/09/23 10:30:42
 * 作者:kerojiang
 */

package wx_api

import "wxapp/model/dto"

type IUserApi interface {

	// 检查加密数据
	//
	// @description:
	//
	// @param:
	//
	// @return:
	CheckEncryptedData(encryptData string) (bool, error)

	// GetOpenId
	// 获取用户的openid
	//  @Description:
	//  @param accessToken
	//  @param code
	//  @return *dto.OpenidResDto
	//  @return error
	//
	GetOpenId(accessToken string, code string) (*dto.OpenidResDto, error)

	// GetPaidUnionId
	// 获取支付后的用户UnionId
	// @description:
	//
	// @param:
	//
	// @return:
	GetPaidUnionId(accessToken string, openId string, transactionId string) (*dto.PayUnionidResDto, error)
}
