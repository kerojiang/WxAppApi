/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:PayRequestDto.go
 * 修改日期:2024/09/23 16:33:09
 * 作者:kerojiang
 */

package dto

import "wxapp/model"

type PayReqDto struct {
	AppId       string       `json:"appId"`
	MchId       string       `json:"mchId"`
	Description string       `json:"description"`
	OutTradeNo  string       `json:"out_trade_no"`
	NotifyUrl   string       `json:"notify_url"`
	Amount      model.Amount `json:"amount"`
	Payer       model.Payer  `json:"payer"`
}
