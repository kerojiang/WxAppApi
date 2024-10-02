/*
 * Copyright (c)  2024  All Rights Reserved
 * 项目名称:WXApp
 * 文件名称:HttpCore.go
 * 修改日期:2024/09/23 10:19:59
 * 作者:kerojiang
 */

package http

import (
	"bytes"
	"io"
	"net/http"

	"wxapp/model/dto"
)

type HttpCore struct {
	config *HttpHeadConfig
}

func NewHttpCore(config *HttpHeadConfig) IHttpCore {

	var _config *HttpHeadConfig
	if config == nil {
		_config = NewHttpHeadConfig()
	} else {
		_config = config
	}

	return &HttpCore{
		config: _config,
	}
}

func (h *HttpCore) Get(url string) (*dto.HttpResDto, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if h.config.Authorization != "" {
		req.Header.Add("Authorization", h.config.Authorization)
	}
	req.Header.Add("Content-Type", h.config.ContentType)
	req.Header.Add("Accept", h.config.Accept)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &dto.HttpResDto{
		StatusCode: resp.StatusCode,
		Body:       string(buf),
	}, nil

}

func (h *HttpCore) Post(url string, data string) (*dto.HttpResDto, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}
	if h.config.Authorization != "" {
		req.Header.Add("Authorization", h.config.Authorization)
	}
	req.Header.Add("Content-Type", h.config.ContentType)
	req.Header.Add("Accept", h.config.Accept)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return &dto.HttpResDto{
		StatusCode: resp.StatusCode,
		Body:       string(buf),
	}, nil
}
