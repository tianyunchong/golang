/**
采集html获取链接
 */
package main

import (
	"caiji/helper"
	"caiji/curl"
	"caiji/preg"
	"fmt"
	"caiji/mysql"
)

func main() {
	/** 初始化mysql链接 */
	config := map[string]string{"user":"root", "pass":"", "host":"127.0.0.1", "port":"3306", "db":"caiji"}
	var cons mysql.Conmysql = mysql.Conmysql{}
	cons.Init(config)
	/** 获取下配置文件 */
	mainConfig := helper.ReadIni("/data/cap/golang/src/caiji/config/config.ini", "main")
	var url string = mainConfig["url"]
	html := curl.GetHtml(url)
	urls := preg.GetUrls(html)
	//往数据库里入mysql
	for _, url := range urls {
		id := mysql.AddUrl(url, cons)
		fmt.Println(id)
	}
}



