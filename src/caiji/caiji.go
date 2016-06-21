/**
采集html获取链接
 */
package main

import (
	"caiji/helper"
	"caiji/curl"
	"caiji/preg"
	"caiji/mysql"
	"fmt"
)

func main() {
	/** 获取下配置文件 */
	mainConfig := helper.ReadIni("/www/gitwork/golang/src/caiji/config/config.ini", "main")
	var url string = mainConfig["url"]
	html := curl.GetHtml(url)
	urls := preg.GetUrls(html)
	//往数据库里入mysql
	for _, url := range urls {
		id := mysql.AddUrl(url)
		fmt.Println(id)
		break
	}
	fmt.Println(urls)
}



