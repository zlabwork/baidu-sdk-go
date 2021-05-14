package request

import (
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "time"
)

type requestData struct {
    reqTime       time.Time
    pubHeader     pubRequestHeader // 公共header
    method        string
    host          string
    uri           string
    header        map[string]string // 补充header
    body          string
    authorization string
}

func (req *requestData) httpRequest() ([]byte, error) {
    url := req.host + req.uri

    // 1. 准备
    client := resty.New()
    request := client.R().
        SetHeaders(req.header).
        // SetAuthToken(req.authorization).
        SetBody([]byte(req.body))

    // 2. 请求
    var resp *resty.Response
    var err error
    if req.method == "GET" {
        resp, err = request.Get(url)
    }
    if req.method == "POST" {
        resp, err = request.Post(url)
    }

    // 3. 结果处理
    if err != nil {
        return nil, err
    }
    if !resp.IsSuccess() {
        return nil, errors.New("request failed: " + url)
    }
    if resp.StatusCode() != 200 {
        return nil, errors.New("response code is " + fmt.Sprintf("%d", resp.StatusCode()))
    }

    return resp.Body(), nil
}
