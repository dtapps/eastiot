package eastiot

import (
	"fmt"
	"go.dtapp.net/gomd5"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"sort"
)

func (c *Client) getSign(p gorequest.Params) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(p.Get(key)))
	}
	signStr += fmt.Sprintf("apiKey=%s", c.GetApiKey())
	return gomd5.ToUpper(signStr)
}
