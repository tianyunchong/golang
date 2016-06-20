/**
* 脚本功能:读取下配置文件
*/
package helper
import (
	"runtime"
	"flag"
	"log"
	"github.com/larspensjo/config"
)
/**
 * 读取配置文件的内容,返回数组
 */
func ReadIni(cpath string, sec string) map[string]string {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	var rs = make(map[string]string)
	configFile := flag.String("configfile", cpath, "General configuration file")
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	if cfg.HasSection(sec) {
		section, err := cfg.SectionOptions(sec)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(sec, v)
				if err == nil {
					rs[v] = options
				}
			}
		}
	}
	return rs
}