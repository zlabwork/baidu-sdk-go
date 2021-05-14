package request

type pubRequestHeader struct {
    Authorization string `json:"authorization"`
    Length        string `json:"content-Length"`
    Type          string `json:"content-Type"`
    Md5           string `json:"content-MD5"`
    Date          string `json:"date"`
    Host          string `json:"host"`
    BceDate       string `json:"x-bce-date"`
}

type pubResponse struct {
    Length     string `json:"Content-Length"`
    Type       string `json:"Content-Type"`
    Md5        string `json:"Content-MD5"`
    Connection string `json:"Connection"`
    Date       string `json:"Date"`
    Tag        string `json:"eTag"`
    Server     string `json:"Server"`
    ReqId      string `json:"x-bce-request-id"`
    DebugId    string `json:"x-bce-debug-id"`
}
