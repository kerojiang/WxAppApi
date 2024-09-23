/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:HttpConfig.go
 * 修改日期:2024/09/23 10:20:27
 * 作者:kerojiang
 */

package http

type HttpConfig struct {
	WechatmpAppid     string `json:"Wechatmp-Appid"`
	WechatmpTimestamp string `json:"Wechatmp-TimeStamp"`
	WechatmpSignature string `json:"Wechatmp-Signature"`
}
