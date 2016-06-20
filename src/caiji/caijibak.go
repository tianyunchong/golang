package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

//指定代理ip
func getTransportFieldURL(proxy_addr *string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(*proxy_addr)
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return
}

//从环境变量$http_proxy或$HTTP_PROXY中获取HTTP代理地址
func getTransportFromEnvironment() (transport *http.Transport) {
	transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	return
}
func fetch(url, proxy_addr *string) (html string) {
	//transport := getTransportFieldURL(proxy_addr)
	//client := &http.Client{Transport: transport}
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	req, err := http.NewRequest("GET", *url, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		html = string(robots)
	} else {
		html = ""
	}
	return
}
func main() {
	proxy_addr := "http://183.221.250.137:80/"
	url := "http://www.gongchang.com/"
	html := fetch(&url, &proxy_addr)
	fmt.Println(html)
}
