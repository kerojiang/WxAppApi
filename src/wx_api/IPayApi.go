/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:IPayApi.go
 * 修改日期:2024/09/23 16:05:36
 * 作者:kerojiang
 */

package wx_api

import "wxapp/model/dto"

type IPayApi interface {

	// 订单支付
	//
	// @description:
	//
	// @param:
	//
	// @return:
	Pay(payRequest *dto.PayReqDto) (*dto.PayResDto, error)

	// 订单查询
	//
	// @description:
	//
	// @param:
	//
	// @return:
	QueryOrder(mchid string, transactionId string) (*dto.QueryOrderResDto, error)

	// 关闭订单
	//
	// @description:
	//
	// @param:
	//
	// @return:
	CloseOrder(closeRequest *dto.CloseOrderReqDto) (bool, error)

	// 订单退款
	//
	// @description:
	//
	// @param:
	//
	// @return:
	Refund(refundRequest *dto.RefundReqDto) (*dto.RefundResDto, error)
}
