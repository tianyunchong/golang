package helper

import (
	"net/url"
	"strings"
	"fmt"
)

/**
	获取完整的url地址
 */
func GetTotalUrl(fromUrl string, curUrl string)string  {
	u, err := url.Parse(curUrl)
	if err != nil {
		fmt.Println(curUrl+"无法解析")
		return ""
	}
	if u.Host != "" {
		return curUrl
	}
	fromU, err := url.Parse(fromUrl)
	if err != nil {
		panic(err)
	}
	sfromU := strings.Split(fromUrl, fromU.Host)
	return sfromU[0]+fromU.Host+curUrl
}
