package mysql

import (
	"strings"
	"fmt"
	"time"
)

/**
 * 增加url地址
 */
func AddUrl(url string, cons Conmysql) int64{
	url = strings.TrimSpace(url)
	url = strings.ToLower(url)
	isExist := ExistUrl(url, cons)
	if isExist {
		fmt.Println(url+"已经存在\n")
		return 0
	}
	//开始插入url
	t := time.Now().Unix()
	tnew := fmt.Sprintf("%d", t)
	id := cons.Insert(map[string]string{"url":url, "addtime":tnew, "uptime":tnew}, "cj_urls")
	return id
}

/**
 *  判断下url在库中是否存在
 */
func ExistUrl(url string, cons Conmysql) bool{
	rs := cons.Findfirst([]string{"url"}, "url = '"+url+"'", "cj_urls")
	if rs != nil{
		return true
	}
	return false
}

