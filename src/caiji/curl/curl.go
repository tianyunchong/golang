/**
* 文件操作获取
*/
package curl

import (
	"net/http"
	"io/ioutil"
)
/**
* 提取html内容
*/
func GetHtml(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}
	body,_ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	bodyStr := string(body)
	return bodyStr
}