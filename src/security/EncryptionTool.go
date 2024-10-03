/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:EncryptionTool.go
 * 修改日期:2024/09/23 11:02:05
 * 作者:kerojiang
 */

package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// sha256加密
//
// @description:
//
// @param:
//
// @return:
func HmacSha256(msg string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(msg))
	return string(hash.Sum(nil))
}

// aes256加密
//
// @description:
//
// @param:
//
// @return:
func Aes256GCM(msg string, secret string) string {

	nonce := genrateNonce()
	hash := md5.Sum([]byte(secret))

	block, err := aes.NewCipher(hash[:])
	if err != nil {
		panic(err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	sealed := aesgcm.Seal(nil, nonce, []byte(msg), nil)
	out := append(nonce, sealed...)

	return hex.EncodeToString(out)
}

func genrateNonce() []byte {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic("生成随机数异常")
	}

	return nonce
}
