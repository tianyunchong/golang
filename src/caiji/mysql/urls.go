package mysql

import (
	//"fmt"
)
import (
	"strings"
	"fmt"
)
/**
 * 增加url地址
 */
func AddUrl(url string) int64{
	url = strings.TrimSpace(url)
	url = strings.ToLower(url)
	isExist := ExistUrl(url)
	if isExist {
		fmt.Println(url+"已经存在\n")
		return 0
	}
	return 1
	//开始插入url
	//con := Conmysql{}
	//id := con.Insert(map[string]string{"url":url}, "cj_urls")
	//return id
}

/**
 *  判断下url在库中是否存在
 */
func ExistUrl(url string) bool{
	config := map[string]string{"user":"root", "pass":"123456", "host":"127.0.0.1", "port":"3306", "db":"caji"}
	con := Conmysql{}
	con.Init(config)
	rs := con.Findfirst([]string{"url"}, "url = '"+url+"'", "cj_urls")
	if rs != nil{
		return true
	}
	return false
}

