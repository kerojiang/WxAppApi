/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:UserApi.go
 * 修改日期:2024/09/23 10:29:37
 * 作者:kerojiang
 */

package wx_api

import (
	"reflect"
	"testing"
	"wxapp/http"
	"wxapp/model/dto"
)

func TestUserApi_GetOpenId(t *testing.T) {
	type fields struct {
		httpCore http.IHttpCore
	}
	type args struct {
		accessToken string
		code        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.OpenidResDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserApi{
				httpCore: tt.fields.httpCore,
			}
			got, err := u.GetOpenId(tt.args.accessToken, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserApi.GetOpenId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserApi.GetOpenId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserApi_GetPaidUnionId(t *testing.T) {
	type fields struct {
		httpCore http.IHttpCore
	}
	type args struct {
		accessToken   string
		openId        string
		transactionId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.PayUnionidResDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserApi{
				httpCore: tt.fields.httpCore,
			}
			got, err := u.GetPaidUnionId(tt.args.accessToken, tt.args.openId, tt.args.transactionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserApi.GetPaidUnionId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserApi.GetPaidUnionId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserApi_GetUserEncryptKey(t *testing.T) {
	type fields struct {
		httpCore http.IHttpCore
	}
	type args struct {
		accessToken string
		openId      string
		sessionKey  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.EncryptKeyResDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserApi{
				httpCore: tt.fields.httpCore,
			}
			got, err := u.GetUserEncryptKey(tt.args.accessToken, tt.args.openId, tt.args.sessionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserApi.GetUserEncryptKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserApi.GetUserEncryptKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserApi_GetPhoneNumber(t *testing.T) {
	type fields struct {
		httpCore http.IHttpCore
	}
	type args struct {
		accessToken string
		phoneCode   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.PhoneInfoResDto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserApi{
				httpCore: tt.fields.httpCore,
			}
			got, err := u.GetPhoneNumber(tt.args.accessToken, tt.args.phoneCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserApi.GetPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserApi.GetPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
