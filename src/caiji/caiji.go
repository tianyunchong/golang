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
	"fmt"
	"sync"
	"strings"
	"strconv"
	"time"
)

var (
	cons mysql.Conmysql
	beansObj helper.BeansCon
	beansObjCheck helper.BeansCon
	beansconfig [2]string
	wg sync.WaitGroup  //定义同步等待组
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/** 初始化mysql链接 */
	config := map[string]string{"user":"root", "pass":"", "host":"127.0.0.1", "port":"3306", "db":"caiji"}
	cons = mysql.Conmysql{}
	cons.Init(config)
	/** 初始化beanstalk连接 */
	beansconfig = [2]string{"tcp", "127.0.0.1:11300"}
	beansObj = helper.BeansCon{}
	beansObj.Init(beansconfig)
	defer beansObj.Close()
	/** 获取下配置文件 */
	mainConfig := helper.ReadIni("/data/cap/golang/src/caiji/config/config.ini", "main")
	beansObj.Write(mainConfig["url"], "go_caiji")
	wg.Add(1)
	go checkUrl()
	wg.Add(1)
	go scanUrls()
	wg.Wait()
	fmt.Println("处理完毕")
}
/**
	读取队列go_check, 检测下url的页面状态
 */
func checkUrl() {
	for {
		url := beansObj.Read("go_check")
		if url == "" {
			fmt.Println("go_check队列里数据已空")
			break
		}
		statusCode := curl.GetCode(url)
		/** 更改下对应的url的状态 */
		cons.Update(map[string]string{"status":fmt.Sprintf("%d", statusCode), "uptime":fmt.Sprintf("%d", time.Now().Unix())}, "url = '"+url+"'", "cj_urls")
	}
	wg.Done()
}

/**
   不停的爬去html页面获取url连接,同时将获取的连接打入队列
 */
func scanUrls() {
	//第一步从队列中获取个连接,然后抓取下url
	for {
		url := beansObj.Read("go_caiji")
		//当队列中没有数据的时候,退出
		if url == "" {
			fmt.Println("抓取url的队列中没有数据了")
			break
		}
		/** 检测下当前连接是否访问过 */
		beginTime := helper.GetDayBeginTime()
		rs := cons.Findfirst([]string{"id", "visittime"}, "url = '"+url+"'", "cj_visit_urls")
		visitTime,_ := strconv.Atoi(rs["visittime"])
		if rs != nil && int64(visitTime) >= beginTime {
			fmt.Println(url+"今天已经爬过了")
			continue
		} else if (rs == nil) {
			cons.Insert(map[string]string{"url":url, "visittime":fmt.Sprintf("%d", beginTime)}, "cj_visit_urls")
		} else {
			//更新下爬取的时间
			cons.Update(map[string]string{"visittime" : fmt.Sprintf("%d", beginTime)}, "id = '"+rs["id"]+"'", "cj_visit_urls")
		}
		html := curl.GetHtml(url)
		urls := preg.GetUrls(html)
		if urls == nil {
			fmt.Println(url+"页面无法抓取到url连接\n")
			continue
		}
		for _, rurl := range urls {
			/** 处理下相对路径 */
			if !strings.Contains(rurl, "/") {
				continue
			}
			rurl = helper.GetTotalUrl(url, rurl)
			if rurl == "" || !strings.Contains(rurl, "gongchang.com") {
				continue
			}
			mysql.AddUrl(map[string]string{"url":rurl, "fromurl":url}, cons)
			beansObj.Write(rurl, "go_caiji")
			beansObj.Write(rurl, "go_check")
		}
	}
	wg.Done()
}




