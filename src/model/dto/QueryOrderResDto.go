/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:QueryOrderResponseDto.go
 * 修改日期:2024/09/23 16:40:58
 * 作者:kerojiang
 */

package dto

type QueryOrderResDto struct {
	MachId         string `json:"mchId"`
	OutTradeNo     string `json:"out_trade_no"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
}
