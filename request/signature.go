package request

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "net/url"
    "sort"
    "strconv"
    "strings"
)

// @link https://cloud.baidu.com/signature/index.html
// @link https://cloud.baidu.com/doc/Reference/s/njwvz1yfu
func (cli *Client) createSignV1() (sign string) {
    // 1.
    ts := cli.request.reqTime
    tsUTC := ts.Format("2006-01-02T15:04:05Z")
    timeout := strconv.FormatInt(cli.timeout, 10)

    // 2. authStringPrefix
    authStringPrefix := "bce-auth-v1/" + cli.credentials.id + "/" + tsUTC + "/" + timeout

    // 3. path escape
    ur, _ := url.Parse(cli.request.uri)
    uriEscape := ur.Path
    if len(ur.RawQuery) > 0 {
        uriEscape += url.QueryEscape("?" + ur.RawQuery)
    }

    // 4. canonicalRequest - 排序、Escape、拼接
    // a>
    body := ""
    if len(cli.request.body) > 0 {
        var keys []string
        for k, _ := range cli.request.body {
            keys = append(keys, k)
        }
        sort.Strings(keys)
        for _, k := range keys {
            body += "&" + k + "=" + url.QueryEscape(cli.request.body[k])
        }
        body = body[1:]
    }

    // b>
    reqRaw := cli.request.method + "\n"
    reqRaw += uriEscape + "\n"
    reqRaw += body + "\n"
    reqRaw += cli.signedHeaders()

    // 5. signingKey
    h := hmac.New(sha256.New, []byte(cli.credentials.secret))
    h.Write([]byte(authStringPrefix))
    signingKey := hex.EncodeToString(h.Sum(nil))

    // 6. signature
    h = hmac.New(sha256.New, []byte(signingKey))
    h.Write([]byte(reqRaw))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature
}

func (cli *Client) createSignV2() {
}

func (cli *Client) signedHeaders() string {
    use := make(map[string]string)
    all := make(map[string]string)
    for k, v := range cli.request.header {
        l := strings.ToLower(k)
        if len(l) >= 5 && l[:5] == "x-bce" {
            use[l] = v
            continue
        }
        all[l] = v
    }

    // 如果出现以下 header 则需要加入到签名
    sh := []string{
        "host",
        "content-length",
        "content-type",
        "content-md5",
    }
    for _, k := range sh {
        _, ok := all[k]
        if ok {
            use[k] = all[k]
        }
    }

    // 排序
    var ks []string
    for k, _ := range use {
        ks = append(ks, k)
    }
    sort.Strings(ks)

    // 拼接
    str := ""
    for _, k := range ks {
        str += k + ":" + url.QueryEscape(use[k]) + "\n"
    }

    return str[:len(str)-1]
}

func (cli *Client) authorizationV1(sign string) string {
    tsUTC := cli.request.reqTime.Format("2006-01-02T15:04:05Z")
    timeout := strconv.FormatInt(cli.timeout, 10)
    return "bce-auth-v1/" + cli.credentials.id + "/" + tsUTC + "/" + timeout + "//" + sign
}
