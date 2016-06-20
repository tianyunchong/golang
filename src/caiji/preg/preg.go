/**
	常用的正则包
 */
package preg

import "regexp"

func GetUrls(str string)[]string {
	reg := regexp.MustCompile(`<a\s+href="(.+)"[^>]+>`)
}
