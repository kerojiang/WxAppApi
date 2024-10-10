/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:UserApi.go
 * 修改日期:2024/09/23 10:29:37
 * 作者:kerojiang
 */

package wx_api

import (
	"errors"

	"wxapp/enum"
	"wxapp/fastjson"
	"wxapp/http"
	"wxapp/model/dto"
	"wxapp/security"

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

func (u *UserApi) GetOpenId(accessToken string, code string) (*dto.OpenidResDto, error) {
	url := "https://api.weixin.qq.com/wxa/getpluginopenpid?access_token=" + accessToken

	reqBody := "{\"code\":\"" + code + "\"}"

	res, err := u.httpCore.Post(url, reqBody)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.OpenidResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, errors.New("获取openid失败")
	}

}

func (u *UserApi) CheckEncryptedData(encryptData string) (bool, error) {
	panic("TODO: Implement")
}

func (u *UserApi) GetPaidUnionId(accessToken string, openId string, transactionId string) (*dto.PayUnionidResDto, error) {
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId + "&transaction_id=" + transactionId

	res, err := u.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {

		resDto, err := fastjson.ConvertToObj[dto.PayUnionidResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, errors.New("获取unionid异常")
	}

}

func (u *UserApi) GetUserEncryptKey(accessToken string, openId string, sessionKey string) (*dto.EncryptKeyResDto, error) {

	signData := security.HmacSha256(sessionKey, "")
	url := "https://api.weixin.qq.com/wxa/business/getuserencryptkey?access_token=" + accessToken + "&signature=" + signData + "&openid=" + openId + "&sig_method=hmac_sha256"

	res, err := u.httpCore.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.EncryptKeyResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, errors.New("获取用户encryptkey失败")
	}
}

func (u *UserApi) GetPhoneNumber(accessToken string, phoneCode string) (*dto.PhoneInfoResDto, error) {

	url := "https://api.weixin.qq.com/wxa/getpluginopenpid?access_token=" + accessToken

	reqBody := "{\"code\":\"" + phoneCode + "\"}"
	res, err := u.httpCore.Post(url, reqBody)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == int(enum.OK) {
		resDto, err := fastjson.ConvertToObj[dto.PhoneInfoResDto](res.Body)
		if err != nil {
			return nil, err
		}
		return resDto, nil
	} else {
		return nil, errors.New("获取用户手机号失败")
	}

}
