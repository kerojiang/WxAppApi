/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:RefundResponseDto.go
 * 修改日期:2024/09/23 16:59:07
 * 作者:kerojiang
 */

package dto

// 退款响应dto
//
// @description:
type RefundResponseDto struct {
	RefundId    string `json:"refund_id"`
	OutRefundNo string `json:"out_refund_no"`
}
