/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:HttpCore.go
 * 修改日期:2024/09/23 10:19:59
 * 作者:kerojiang
 */

package http

import "testing"

func TestHttpCore_Get(t *testing.T) {
	url := "https://www.baidu.com"
	httpCore := NewHttpCore(nil)

	result, err := httpCore.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fail()
	}

}
