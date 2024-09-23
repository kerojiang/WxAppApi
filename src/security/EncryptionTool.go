/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:EncryptionTool.go
 * 修改日期:2024/09/23 11:02:05
 * 作者:kerojiang
 */

package security

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(msg string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))
	return string(h.Sum(nil))
}
