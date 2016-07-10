/**
 * 写文件
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	userfile := "test.txt"
	fout, err := os.Create(userfile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("jtesttsst\n")
		fout.Write([]byte("test asdjflsfjdlsafd\n"))
	}
}
