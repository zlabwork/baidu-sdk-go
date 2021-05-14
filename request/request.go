package request

import (
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

func (req *requestData) httpRequest() {
}
