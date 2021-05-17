package request

import (
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "time"
)

type requestData struct {
    reqTime time.Time
    method  string
    host    string
    uri     string
    header  map[string]string
    body    string
}

func (req *requestData) httpRequest() ([]byte, error) {
    requestUrl := "https://" + req.host + req.uri

    // 1. 准备
    client := resty.New()
    client.SetDebug(true) // TODO :: http debug
    client.SetContentLength(true)

    // request
    request := client.R().
        SetHeaders(req.header).
        SetBody([]byte(req.body))

    // 2. 请求
    var resp *resty.Response
    var err error
    if req.method == "GET" {
        resp, err = request.Get(requestUrl)
    }
    if req.method == "POST" {
        resp, err = request.Post(requestUrl)
    }
    if req.method == "HEAD" {
        resp, err = request.Head(requestUrl)
    }

    // 3. 结果处理
    if err != nil {
        return nil, err
    }
    if !resp.IsSuccess() {
        return nil, errors.New("request failed: " + requestUrl)
    }
    if resp.StatusCode() != 200 {
        return nil, errors.New("response code is " + fmt.Sprintf("%d", resp.StatusCode()))
    }

    return resp.Body(), nil
}
