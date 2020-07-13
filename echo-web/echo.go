package main

import (
	"flag"

	"letinvr.com/echo-web/module/log"
	"letinvr.com/echo-web/router"
)

const (
	DefaultConfFilePath = "conf/conf.toml"
)

var (
	confFilePath string
	cmdHelp      bool
)

func init() {
	flag.StringVar(&confFilePath, "c", DefaultConfFilePath, "配置文件路径")
	flag.BoolVar(&cmdHelp, "h", false, "帮助")
	flag.Parse()
}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}
	log.Debugf("run with conf:%s", confFilePath)

	// 子域名部署
	router.RunSubdomains(confFilePath)
}
