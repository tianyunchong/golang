/**
采集html获取链接
 */
package main

import (
	"caiji/helper"
	"caiji/mysql"
	"runtime"
	"caiji/curl"
	"caiji/preg"
	"time"
	"fmt"
	"sync"
)

var (
	cons mysql.Conmysql
	beansObj helper.BeansCon
	wg sync.WaitGroup  //定义同步等待组
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/** 初始化mysql链接 */
	config := map[string]string{"user":"root", "pass":"", "host":"127.0.0.1", "port":"3306", "db":"caiji"}
	cons = mysql.Conmysql{}
	cons.Init(config)
	/** 初始化beanstalk连接 */
	var beansconfig = [2]string{"tcp", "127.0.0.1:11300"}
	beansObj = helper.BeansCon{}
	beansObj.Init(beansconfig, "go_caiji")
	defer beansObj.Close()
	/** 获取下配置文件 */
	mainConfig := helper.ReadIni("/data/cap/golang/src/caiji/config/config.ini", "main")
	beansObj.Write(mainConfig["url"])
	wg.Add(1)
	go scanUrls()
	wg.Wait()
}

/**
   不停的爬去html页面获取url连接,同时将获取的连接打入队列
 */
func scanUrls() {
	//第一步从队列中获取个连接,然后抓取下url
	for {
		url := beansObj.Read()
		if url == "" {
			time.Sleep(10 * time.Second)
			continue
		}
		html := curl.GetHtml(url)
		urls := preg.GetUrls(html)
		if urls == nil {
			fmt.Println(url+"页面无法抓取到url连接\n")
			continue
		}
		for _, url := range urls {
			mysql.AddUrl(url, cons)
			beansObj.Write(url)
		}
	}
	wg.Done()
}




