/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:RefundRequestDto.go
 * 修改日期:2024/09/23 16:58:54
 * 作者:kerojiang
 */

package dto

import "wxapp/model"

// 退款请求dto
//
// @description:
type RefundRequestDto struct {
	OutRefundNo string       `json:"out_refund_no"`
	Amount      model.Amount `json:"amount"`
}
