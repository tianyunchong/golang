/**
	常用的正则包
 */
package preg

import (
	"regexp"
)

/**
   从字符串中提取出所有的url
 */
func GetUrls(str string)[]string {
	var rs []string;
	reg := regexp.MustCompile(`<a\s+href="([^"']+)"[^>]+>`)
	locArr := reg.FindAllStringSubmatchIndex(str, -1)
	for _, loc := range locArr {
		//fmt.Println(loc)
		rs = append(rs, str[loc[2] : loc[3]])
	}
	return rs
}
