/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:AppBaseApi.go
 * 修改日期:2024/09/23 10:29:27
 * 作者:kerojiang
 */

package wx_api

import (
	"errors"

	"wxapp/enum"
	"wxapp/fastjson"
	"wxapp/http"
	"wxapp/security"

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

func (a *AppBaseApi) GetAccessToken(appId string, secret string) (*dto.AccessTokenResDto, error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + secret
	res, err := a.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.AccessTokenResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, err
	}
}

func (a *AppBaseApi) Code2Session(appId string, secret string, code string) (*dto.Code2SessionResDto, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code"
	res, err := a.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.Code2SessionResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, err
	}
}

func (a *AppBaseApi) CheckSessionKey(accessToken string, sessionKey string, openid string) (bool, error) {

	signData := security.HmacSha256(sessionKey, "")

	url := "https://api.weixin.qq.com/wxa/checksession?access_token=" + accessToken + "&signature=" + signData + "&openid=" + openid + "&sig_method=hmac_sha256"
	res, err := a.httpCore.Get(url)
	if err != nil {
		return false, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.WXAppErrorDto](res.Body)
		if err != nil {
			return false, err
		}
		if resDto.ErrCode == 0 && resDto.ErrMsg == "ok" {
			return true, nil
		}
		return false, errors.New(resDto.ErrMsg)

	} else {
		return false, errors.New("检查登陆状态失败")
	}
}

func (a *AppBaseApi) ResetUserSessionKey(accessToken string, sessionKey string,
	openid string) (*dto.Code2SessionResDto, error) {

	signData := security.HmacSha256(sessionKey, "")

	url := "https://api.weixin.qq.com/wxa/resetusersessionkey?access_token=" + accessToken + "&openid=" + openid + "&signature=" + signData + "&sig_method=hmac_sha256"

	res, err := a.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.Code2SessionResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, errors.New("重置登陆状态失败")
	}

}
