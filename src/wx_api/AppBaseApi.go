/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:AppBaseApi.go
 * 修改日期:2024/09/23 10:29:27
 * 作者:kerojiang
 */

package wx_api

import (
	"wxapp/http"

	"wxapp/model/dto"

	"github.com/samber/do/v2"
)

type AppBaseApi struct {
	httpCore http.IHttpCore
}

func NewAppBaseApi(i do.Injector) IAppBaseApi {
	return &AppBaseApi{
		httpCore: do.MustInvoke[http.IHttpCore](i),
	}
}

func (a *AppBaseApi) GetAccessToken(appId string, secret string) (*dto.AccessTokenDto, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + secret
	result, err := a.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
}

func (a *AppBaseApi) Code2Session(appId string, secret string, code string) (*dto.Code2SessionDto, error) {
	panic("TODO: Implement")
}

func (a *AppBaseApi) CheckSessionKey(accessToken string, openid string) (bool, error) {
	panic("TODO: Implement")
}

func (a *AppBaseApi) ResetUserSessionKey(accessToken string, openid string) (*dto.Code2SessionDto, error) {
	panic("TODO: Implement")
}
