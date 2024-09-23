/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:ErrorCode.go
 * 修改日期:2024/09/23 10:46:25
 * 作者:kerojiang
 */

package enum

// 微信小程序错误码
type ErrorCode int

const (
	SystemError       ErrorCode = -1
	CodeBlocked       ErrorCode = 40226
	ApiMustSlower     ErrorCode = 45011
	InvalidCode       ErrorCode = 40029
	InvalidSigmethod  ErrorCode = 87008
	SessionkeyExpired ErrorCode = 87007
	InvalidSignature  ErrorCode = 87009
	InvalidArgs       ErrorCode = 40097
)
