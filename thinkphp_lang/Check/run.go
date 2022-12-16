package Check

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"thinkphp_lang/Common"
	"time"
)

func Check_url(Cmd_instruction Common.Cmd) {

	if Cmd_instruction.Target != "" && Cmd_instruction.Target != "12" {
		fmt.Println("开始扫描...")
		startTime := time.Now()
		Common.CustomizeGET(AnalyzeUrl(Cmd_instruction.Target))
		//fmt.Println(result)
		endTime := time.Since(startTime)
		fmt.Printf("\n扫描结束,耗时: %v\n", endTime)

	}
	// 批量扫描
	if Cmd_instruction.Targets != "" {
		startTime := time.Now()

		fmt.Println("开始批量扫描...")
		Targets := Readfiles(Cmd_instruction.Targets)
		var wg sync.WaitGroup
		var taskChan = make(chan string, len(Targets))
		for _, url := range Targets {
			taskChan <- url

		}
		close(taskChan)

		for i := 0; i <= Cmd_instruction.Thread; i++ {
			wg.Add(1)
			// v := <-taskChan
			go func() {
				for {
					target, ok := <-taskChan
					//fmt.Println(target)
					//fmt.Println(ok)
					if !ok {

						break
					}
					Common.CustomizeGET(AnalyzeUrl(target))

				}
				wg.Done()
			}()

		}
		wg.Wait()
		WriteFile(Cmd_instruction.Output, Common.ResultUrl)
		endTime := time.Since(startTime)
		fmt.Printf("\n[*]扫描结束，共耗时: %v\n", endTime)
		fmt.Printf("\n[*]扫描结果保存至%v中", Cmd_instruction.Output)

	}
	// 上传shell
	if Cmd_instruction.Webshell != "" && Cmd_instruction.Target == "" && Cmd_instruction.Targets == "" {
		startTime := time.Now()
		fmt.Println("开始上传webshell...")
		Common.ExpGet(Cmd_instruction.Webshell)
		endTime := time.Since(startTime)
		fmt.Printf("\n[*]扫描结束，共耗时: %v\n", endTime)
	}

}

func WriteFile(OutPutName string, data []string) {
	if !strings.HasSuffix(OutPutName, ".txt") {
		OutPutName = OutPutName + ".txt"
	}
	f, _ := os.OpenFile(OutPutName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	write := bufio.NewWriter(f)
	for _, r := range data {
		write.WriteString(r)

	}
	write.Flush()

}

// 读取文件
func Readfiles(FileName string) []string {
	var readcontent []string
	r, _ := os.Open(FileName)
	defer r.Close()
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		readcontent = append(readcontent, line)
	}

	return readcontent
}

// 解析url
func AnalyzeUrl(Url string) string {
	u, err := url.Parse(Url)
	if err != nil {
		return ""
	}
	if u.Scheme != "" {
		return u.Scheme + "://" + u.Host
	} else {
		if Common.RequestGET("http://"+u.Host) != "200" {
			return "https://" + u.Host
		} else {
			return "http://" + u.Host
		}

	}

}
