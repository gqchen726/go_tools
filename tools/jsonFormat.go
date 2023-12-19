package tools

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func JsonFormatFromConsole(jsonStr string) string {
	// 使用正则表达式判断字符传的加密类型是否为base64
	regex := regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`)
	if regex.MatchString(jsonStr) {
		fmt.Println("字符串符合Base64编码规则")
		// base64解码
		decodeBytes, err := base64.StdEncoding.DecodeString(jsonStr)
		if err != nil {
			fmt.Println("base64解码失败:", err)
		}
		jsonStr = string(decodeBytes)
	}
	// 处理字符串中可能有的base64字符串片段
	base64SubStringRegex := regexp.MustCompile(`[A-Za-z0-9+/]+={1,2}`)
	encodedStrings := base64SubStringRegex.FindAllString(jsonStr, -1)
	// 遍历encodedStrings
	for _, encodedString := range encodedStrings {
		// base64解码
		decodeBytes, err := base64.StdEncoding.DecodeString(encodedString)
		if err == nil {
			// base64解码后的字符串
			jsonStrByDecode := string(decodeBytes)
			jsonStrByDecode = strings.ReplaceAll(jsonStrByDecode, `"`, `\"`)
			jsonStrByDecode = strings.ReplaceAll(jsonStrByDecode, `{`, `"{`)
			// jsonStrByDecode = strings.ReplaceAll(jsonStrByDecode, `}`, `}"`)
			jsonStrByDecode = strings.ReplaceAll(jsonStrByDecode, `[`, `"[`)
			jsonStrByDecode = strings.ReplaceAll(jsonStrByDecode, `]`, `]"`)
			// 替换字符串
			// fmt.Println("encodedString: ", encodedString, "jsonStrByDecode: ", jsonStrByDecode)
			if utf8.ValidString(jsonStrByDecode) {
				jsonStr = strings.ReplaceAll(jsonStr, encodedString, jsonStrByDecode)
			}
		}
	}
	s := strings.ReplaceAll(jsonStr, `\"`, `"`)
	s = strings.ReplaceAll(s, `"{`, `{`)
	s = strings.ReplaceAll(s, `}"`, `}`)
	s = strings.ReplaceAll(s, `"{"`, `{"`)
	s = strings.ReplaceAll(s, `"}"`, `"}`)
	s = strings.ReplaceAll(s, `"[{`, `[{`)
	s = strings.ReplaceAll(s, `}]"`, `}]`)
	return s
}

func JsonFormatFromConsoleForPrettyJson(jsonStr string) []byte {
	// 使用正则表达式校验jsonStr是否为json格式
	regex := regexp.MustCompile(`^\s*{.*}\s*$`)
	if !regex.MatchString(jsonStr) {
		fmt.Println("不是json格式字符串")
		return nil
	}
	// 使用json格式化json格式数据s
	s := JsonFormatFromConsole(jsonStr)
	// 解析JSON字符串
	var data interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		fmt.Println("JSONString:", s)
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
	// 读取文件
	file, err := os.Open(inFilePath)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	outputString := ""
	for {
		line, err := reader.ReadString('\n')

		// 判断换行符
		switch {
		case err == nil:
			// 行末尾有换行符
			// fmt.Println("行:", line)
			outputString = readLineAndHandle(outputString, line)
		case err.Error() == "EOF":
			// 文件结束，最后一行可能没有换行符
			// fmt.Println("行:", line)
			outputString = readLineAndHandle(outputString, line)
			os.WriteFile(outFilePath, []byte(outputString), 0644)
			// fmt.Println("文件结束")
			return
		default:
			// 处理读取错误
			fmt.Println("读取错误:", err)
			return
		}
	}
}

func readLineAndHandle(outputString string, line string) string {
	if line == "" {
		return outputString
	}
	prettyJSON := JsonFormatFromConsoleForPrettyJson(line)
	// 将s写入outFilePath文件
	outputString += string(prettyJSON)
	outputString += "\n"
	outputString += "\n"
	return outputString
}
