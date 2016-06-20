/**
采集html获取链接
 */
package main

import (
	"caiji/helper"
	"fmt"
	"caiji/curl"
)

func main() {
	/** 获取下配置文件 */
	mainConfig := helper.ReadIni("/www/gitwork/golang/src/caiji/config/config.ini", "main")
	var url string = mainConfig["url"]
	html := curl.GetHtml(url)
	fmt.Println(html)
}
