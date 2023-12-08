package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func JsonFormatFromConsole(jsonStr string) string {
	s := strings.ReplaceAll(jsonStr, `\"`, `"`)
	s = strings.ReplaceAll(s, `"{"`, `{"`)
	s = strings.ReplaceAll(s, `"}"`, `"}`)
	return s
}

func JsonFormatFromConsoleForPrettyJson(jsonStr string) []byte {
	// 使用json格式化json格式数据s
	s := JsonFormatFromConsole(jsonStr)
	// 解析JSON字符串
	var data interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
	}

	// 格式化JSON
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("格式化JSON失败:", err)
	}
	return prettyJSON
}

func JsonFormatFromConsoleForPrettyJsonString(jsonStr string) string {
	return string(JsonFormatFromConsoleForPrettyJson(jsonStr))
}

func JsonFormatFromFile(inFilePath string, outFilePath string) {
	if inFilePath == "" {
		inFilePath = "C:\\Temp\\test-in.txt"
	}
	if outFilePath == "" {
		outFilePath = "C:\\Temp\\test-out.txt"
	}
	// 读取filePath文件的第一行内容
	jsonStr := ""
	// 判断inFilePath文件是否存在,如果不存在则创建
	if _, err := os.Stat(inFilePath); os.IsNotExist(err) {
		_, err := os.Create(inFilePath)
		if err != nil {
			panic(err)
		}
	}
	// 判断outFilePath文件是否存在,如果不存在则创建
	if _, err := os.Stat(outFilePath); os.IsNotExist(err) {
		_, err := os.Create(outFilePath)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.ReadFile(inFilePath)
	// handle err
	if err != nil {
		panic(err)
	}
	// 将f转为string赋值给jsonStr
	jsonStr = string(f)
	prettyJSON := JsonFormatFromConsoleForPrettyJson(jsonStr)
	// 将s写入outFilePath文件
	os.WriteFile(outFilePath, []byte(prettyJSON), 0644)
}
