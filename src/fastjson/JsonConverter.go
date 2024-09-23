/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:JsonConverter.go
 * 修改日期:2024/09/23 14:57:41
 * 作者:kerojiang
 */

package fastjson

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 序列化为json
//
// @description:
//
// @param:
//
// @return:
func ConvertToStr[T any](data *T) (string, error) {

	buf, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// 反序列化为struct
//
// @description:
//
// @param:
//
// @return:
func ConvertToObj[T any](data string) (*T, error) {
	var obj *T
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
