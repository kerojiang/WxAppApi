/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:PayApi.go
 * 修改日期:2024/09/23 16:05:44
 * 作者:kerojiang
 */

package wx_api

import (
	"wxapp/http"
	"wxapp/model/dto"

	"github.com/samber/do/v2"
)

type PayApi struct {
	httpCore http.IHttpCore
}

func NewPayApi(i do.Injector) IPayApi {
	return &PayApi{
		httpCore: do.MustInvoke[http.IHttpCore](i),
	}

}

func (p *PayApi) Pay(payRequest *dto.PayReqDto) (*dto.PayResDto, error) {
	panic("TODO: Implement")
}

func (p *PayApi) QueryOrder(mchid string, transactionId string) (*dto.QueryOrderResDto, error) {
	panic("TODO: Implement")
}

func (p *PayApi) CloseOrder(closeRequest *dto.CloseOrderReqDto) (bool, error) {
	panic("TODO: Implement")
}

func (p *PayApi) Refund(refundRequest *dto.RefundReqDto) (*dto.RefundResDto, error) {
	panic("TODO: Implement")
}
