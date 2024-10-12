/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:JsonConverter.go
 * 修改日期:2024/09/23 14:57:41
 * 作者:kerojiang
 */

package fastjson

import (
	"fmt"
	"testing"
	"wxapp/model"
	"wxapp/model/dto"
)

func TestConvertToStr(t *testing.T) {
	type TestA struct {
		Name string
		Age  int
	}

	tmodel := &TestA{
		Name: "test",
		Age:  5,
	}
	result, err := ConvertToStr[TestA](tmodel)
	if err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Fail()
	}

	data, err := ConvertToObj[TestA](result)
	if err != nil {
		t.Fatal(err)
	}

	if data.Name != "test" {
		t.Fail()
	}

	keys := []*model.EncryptKey{}

	for i := 0; i < 5; i++ {
		encryKey := &model.EncryptKey{
			Version:    1,
			ExpireIn:   7200,
			EncryptKey: "test",
			Iv:         "test",
			CreateTime: 123456789,
		}
		keys = append(keys, encryKey)
	}

	encryResDto := &dto.EncryptKeyResDto{
		WXAppErrorDto: &dto.WXAppErrorDto{
			ErrMsg:  "test",
			ErrCode: 10004,
		},
		KeyInfoList: keys,
	}

	str, err := ConvertToStr(encryResDto)
	if err != nil {
		t.Fatal(err)
	}
	resDto, err := ConvertToObj[dto.EncryptKeyResDto](str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resDto)

}
