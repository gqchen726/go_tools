package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 定义目标网站的 URL
const targetURL = "http://www.example.com/signin"

// 定义请求头
var headers = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
}

// 定义签到成功的标志
var checkinedFlag = false

// 定义主函数
func main() {
	// 获取当前日期
	now := time.Now()
	year, month, day := now.Date()

	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 HTTP GET 请求
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取 HTTP 响应正文
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 检查 HTTP 响应状态码
	if resp.StatusCode != 200 {
		log.Fatalf("HTTP 请求失败，状态码：%d", resp.StatusCode)
	}

	// 检查是否已经签到
	if strings.Contains(string(body), "您已经签到") {
		fmt.Printf("%d年%d月%d日，您已经签到。\n", year, month, day)
		checkinedFlag = true
	}

	// 如果没有签到，则进行签到
	if !checkinedFlag {
		// 获取签到按钮的表单数据
		formValues := map[string][]string{
			"签到": {"签到"},
		}

		// 创建一个 HTTP POST 请求
		req, err = http.NewRequest("POST", targetURL, strings.NewReader(encodeFormValues(formValues)))
		if err != nil {
			log.Fatal(err)
		}

		// 设置请求头
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		// 设置请求正文的 Content-Type
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// 发送 HTTP 请求
		resp, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// 读取 HTTP 响应正文
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 检查 HTTP 响应状态码
		if resp.StatusCode != 200 {
			log.Fatalf("HTTP 请求失败，状态码：%d", resp.StatusCode)
		}

		// 检查是否签到成功
		if strings.Contains(string(body), "checkinedFlag") {
			fmt.Printf("%d年%d月%d日，签到成功。\n", year, month, day)
			checkinedFlag = true
		} else {
			fmt.Printf("%d年%d月%d日，签到失败。\n", year, month, day)
		}
	}
}

// 将表单数据编码为字符串
func encodeFormValues(formValues map[string][]string) string {
	var payload strings.Builder
	for key, values := range formValues {
		for _, value := range values {
			payload.WriteString(strconv.Quote(key))
			payload.WriteString("=")
			payload.WriteString(strconv.Quote(value))
			payload.WriteString("&")
		}
	}
	return payload.String()
}
