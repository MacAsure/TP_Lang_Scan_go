package Common

import (
	"flag"
	"fmt"
)

type Cmd struct {
	Target   string
	Targets  string
	Thread   int
	Output   string
	Webshell string
}

const Banner = `
	
▄▄▄█████▓ ██▓███   ██▓    ▄▄▄       ███▄    █   ▄████   ██████  ▄████▄   ▄▄▄       ███▄    █ 
▓  ██▒ ▓▒▓██░  ██▒▓██▒   ▒████▄     ██ ▀█   █  ██▒ ▀█▒▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █ 
▒ ▓██░ ▒░▓██░ ██▓▒▒██░   ▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒
░ ▓██▓ ░ ▒██▄█▓▒ ▒▒██░   ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓  ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒
  ▒██▒ ░ ▒██▒ ░  ░░██████▒▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░
  ▒ ░░   ▒▓▒░ ░  ░░ ▒░▓  ░▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
    ░    ░▒ ░     ░ ░ ▒  ░ ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ ░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░
  ░      ░░         ░ ░    ░   ▒      ░   ░ ░ ░ ░   ░ ░  ░  ░  ░          ░   ▒      ░   ░ ░ 
                      ░  ░     ░  ░         ░       ░       ░  ░ ░            ░  ░         ░ 
                                                               ░                             
																	      						--code-by: iceberg-N
开发人员不承担任何责任,也不对任何滥用或者损坏负责
此工具为thinkphp多语言rce综合利用工具,利用仅针对docker环境

`

func Flag(info *Cmd) {
	fmt.Println(Banner)
	flag.StringVar(&info.Target, "u", "", "指定目标url")
	flag.StringVar(&info.Targets, "f", "", "指定文件路径")
	flag.StringVar(&info.Output, "o", "result", "默认保存为result")
	flag.IntVar(&info.Thread, "t", 10, "设置线程,默认10")
	flag.StringVar(&info.Webshell, "w", "", "上转webshell,输入想上传webshell的url")

	flag.Parse()
}
