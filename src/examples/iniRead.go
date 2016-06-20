// ini config read
package main

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"log"
	"runtime"
)

var (
	cpath = "/www/gitwork/golang/src/examples/config/config.ini"
	configFile = flag.String("configfile", cpath, "General configuration file")
)

var TOPIC = make(map[string]string)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	if cfg.HasSection("topicArr") {
		section, err := cfg.SectionOptions("topicArr")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	fmt.Println(TOPIC)
	fmt.Println(TOPIC["addr"])
	fmt.Println(TOPIC["debug"])
}
