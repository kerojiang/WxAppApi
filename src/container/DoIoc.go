/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:src
 * 文件名称:DoInstance.go
 * 修改日期:2024/09/23 11:51:09
 * 作者:kerojiang
 */

package container

import (
	"sync"

	"github.com/samber/do/v2"
)

var once sync.Once

var injector *do.RootScope

// ioc容器实例
//
// @description:
//
// @param:
//
// @return:
func DoInstance() *do.RootScope {

	once.Do(func() {
		injector = do.New()
	})
	return injector
}
