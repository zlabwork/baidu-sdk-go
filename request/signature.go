package request

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "net/url"
    "sort"
    "strconv"
    "time"
)

// @link https://cloud.baidu.com/signature/index.html
// @link https://cloud.baidu.com/doc/Reference/s/njwvz1yfu
func (cli *Client) createSignV1(method string, uri string) (sign string, timestamp time.Time) {
    // 1.
    ts := time.Now()
    tsUTC := ts.Format("2006-01-02T15:04:05Z")
    timeout := strconv.FormatInt(cli.timeout, 10)

    // 2. authStringPrefix
    authStringPrefix := "bce-auth-v1/" + cli.credentials.id + "/" + tsUTC + "/" + timeout

    // 3. canonicalRequest - 排序、Escape、拼接
    // a>
    ur, _ := url.Parse(uri)
    var keys []string
    for k, _ := range ur.Query() {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    query := ""
    for _, k := range keys {
        query += "&" + k + "=" + url.QueryEscape(ur.Query()[k][0])
    }
    query = query[1:]
    // b>
    reqRaw := method + "\n"
    reqRaw += ur.Path + "\n"
    reqRaw += query + "\n"
    reqRaw += "host:" + cli.host + "\n"
    reqRaw += "x-bce-date:" + url.QueryEscape(tsUTC)

    // 4. signingKey
    h := hmac.New(sha256.New, []byte(cli.credentials.secret))
    h.Write([]byte(authStringPrefix))
    signingKey := hex.EncodeToString(h.Sum(nil))

    // 5. signature
    h = hmac.New(sha256.New, []byte(signingKey))
    h.Write([]byte(reqRaw))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature, ts
}

func (cli *Client) createSignV2() {
}

func (cli *Client) authorizationV1(sign string, ts time.Time) string {
    tsUTC := ts.Format("2006-01-02T15:04:05Z")
    timeout := strconv.FormatInt(cli.timeout, 10)
    return "bce-auth-v1/" + cli.credentials.id + "/" + tsUTC + "/" + timeout + "/host;x-bce-date/" + sign
}
