package main

import (
    	"fmt"
	"net/url"
	//"strings"
	"strings"
)

func main() {
	fmt.Println(GetTotalUrl("http://www.baidu.com", "http://al.job1001.com"))
	//var surl string = "/test/?age=11"
    	//u, err := url.Parse(surl)
    	//if err != nil {
        	//panic(err)
    	//}
	//if u.Host == "" {
	//	fmt.Println("tests")
	//}
	//fmt.Println(u.Host)
	//h := strings.Split(surl, u.Host)
	//fmt.Println(h[0]+u.Host)
}


func GetTotalUrl(fromUrl string, curUrl string)string  {
	u, err := url.Parse(curUrl)
	if err != nil {
		panic(err)
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