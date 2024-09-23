/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:JsonConverter.go
 * 修改日期:2024/09/23 14:57:41
 * 作者:kerojiang
 */

package fastjson

import (
	"testing"
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

}
