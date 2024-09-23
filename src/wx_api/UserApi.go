/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:UserApi.go
 * 修改日期:2024/09/23 10:29:37
 * 作者:kerojiang
 */

package wx_api

import (
	"wxapp/http"

	"github.com/samber/do/v2"
)

type UserApi struct {
	httpCore http.IHttpCore
}

func NewUserApi(i do.Injector) IUserApi {
	return &UserApi{
		httpCore: do.MustInvoke[http.IHttpCore](i),
	}
}

func (u *UserApi) CheckEncryptedData(encryptData string) (bool, error) {
	panic("TODO: Implement")
}

func (u *UserApi) GetPaidUnionId(openId string, transactionId string) (string, error) {
	panic("TODO: Implement")
}
