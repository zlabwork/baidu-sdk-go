package vision

import (
    "encoding/base64"
    "github.com/zlabwork/baidu-sdk-go/request"
    "io/ioutil"
    "os"
)

const (
    endpoint = "aip.baidubce.com"
)

type vision struct {
    cli *request.Client
}

func NewVision() *vision {
    v := &vision{
        cli: request.NewClient(),
    }
    v.cli.SetHost(endpoint)
    return v
}

// 场景检测
// @link https://ai.baidu.com/ai-doc/IMAGERECOGNITION/Xk3bcxe21
func (vi *vision) Scene(address string) ([]byte, error) {

    // 1. uri
    uri := "/rest/2.0/image-classify/v2/advanced_general"

    // 2. header
    header := make(map[string]string)
    header["Content-Type"] = "application/x-www-form-urlencoded"

    // 3. body
    body := make(map[string]string)
    body["baike_num"] = "0"
    if address[:4] == "http" {
        body["url"] = address
    } else {
        f, err := os.Open(address)
        if err != nil {
            return nil, err
        }
        defer f.Close()
        b, err := ioutil.ReadAll(f)
        if err != nil {
            return nil, err
        }
        encode := base64.StdEncoding
        body["image"] = encode.EncodeToString(b)
    }

    // 4. result
    resp, err := vi.cli.BuildRequest("POST", uri, header, body)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

// 地标识别
// @link https://ai.baidu.com/ai-doc/IMAGERECOGNITION/jk3bcxbih
func (vi *vision) Landmark(address string) ([]byte, error) {

    // 1. uri
    uri := "/rest/2.0/image-classify/v1/landmark"

    // 2. header
    header := make(map[string]string)
    header["Content-Type"] = "application/x-www-form-urlencoded"

    // 3. body
    body := make(map[string]string)
    if address[:4] == "http" {
        body["url"] = address
    } else {
        f, err := os.Open(address)
        if err != nil {
            return nil, err
        }
        defer f.Close()
        b, err := ioutil.ReadAll(f)
        if err != nil {
            return nil, err
        }
        encode := base64.StdEncoding
        body["image"] = encode.EncodeToString(b)
    }

    // 4. result
    resp, err := vi.cli.BuildRequest("POST", uri, header, body)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
