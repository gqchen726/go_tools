package consoles

import (
	"fmt"
	"go_tools/tools"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

func JsonFormatFromConsole(input string) bool {
	tools.ClearConsole()
	// 读取用户home路径
	userHome := os.Getenv("USERPROFILE")
	var basePath string = userHome + "\\Documents"
	var inFilePath string = basePath + "\\json-in.json"
	var outFilePath string = basePath + "\\json-out.json"
	if input == "jof" {
		switch tools.ReturnMonitoringKeyboard() {
		case keyboard.KeyEsc:
			return true
		case keyboard.KeyBackspace:
			return false
		}
		tools.JsonFormatFromFile(inFilePath, outFilePath)
	} else {
		for {
			fmt.Println("------------------------------")
			fmt.Printf("%s\n%s\n%s\n", "支持的操作(请输入编号):", "1. from console", "2. from file")
			switch tools.ReturnMonitoringKeyboard() {
			case keyboard.KeyEsc:
				return true
			case keyboard.KeyBackspace:
				return false
			case keyboard.KeyEnter:
				fmt.Println("请输入命令编号：")
			}
			var operation string
			fmt.Scanln(&operation)
			switch operation {
			case "2":
				var inFilePathFromInput string
				var outFilePathFromInput string
				fmt.Fscanln(os.Stdin, &inFilePathFromInput)
				fmt.Fscanln(os.Stdin, &outFilePathFromInput)
				if inFilePathFromInput != "" {
					inFilePath = inFilePathFromInput
				}
				if outFilePathFromInput != "" {
					outFilePath = outFilePathFromInput
				}
				fmt.Println("输入的filePath:", inFilePath, "输出的filePath:", outFilePath)
				tools.JsonFormatFromFile(inFilePath, outFilePath)
			case "1":
			default:
				// 接收输入
				fmt.Println("请输入JSON字符串:")
				var jsonStr string
				fmt.Fscanln(os.Stdin, &jsonStr)
				// 替换jsonStr字符串的空格符
				jsonStr = strings.ReplaceAll(jsonStr, " ", "")
				fmt.Println("输入的JSON字符串:", jsonStr)
				fmt.Println(tools.JsonFormatFromConsoleForPrettyJsonString(jsonStr))
			}
			switch tools.ReturnMonitoringKeyboard() {
			case keyboard.KeyEsc:
				return true
			case keyboard.KeyBackspace:
				return false
			case keyboard.KeyEnter:
			}
		}
	}
	return false
}
