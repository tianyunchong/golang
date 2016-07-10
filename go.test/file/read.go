/**
 * 文件操作
 * 读取
 */
package main

import (
	"bufio"
	"code.google.com/p/mahonia"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	filepath := "./log_" + strings.Replace(t.String()[:19], ":", "_", 3) + ".txt"
	fp, err := os.OpenFile(filepath, os.O_CREATE, 0666)
	defer fp.Close()
	if err != nil {
		fmt.Println("文件test.txt", err)
	}
	wFile := bufio.NewWriter(fp)
	wFile.WriteString(readfile())
	wFile.Flush()
}

func readfile() string {
	f, err := os.Open("c:/test/test.txt")
	if err != nil {
		return err.Error()
	}
	defer f.Close()
	buf := make([]byte, 1024)
	//文件ex7.txt的编码是gb18030
	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		return "编码不存在!"
	}
	var str string = ""
	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		}
		//解码为UTF-8
		str += decoder.ConvertString(string(buf[:n]))
	}
	fmt.Println(str)
	return str
}
