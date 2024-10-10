/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WxAppApi
 * 文件名称:HttpStatusCode.go
 * 修改日期:2024/10/10 21:31:55
 * 作者:kerojiang
 */

package enum

type HttpStatusCode int

const (
	//成功
	OK HttpStatusCode = 200
	//请求错误
	BadRequest HttpStatusCode = 400
	//内部服务错误
	InternalServerError HttpStatusCode = 500
)
