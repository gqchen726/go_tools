package main

import (
	"fmt"
	"go_tools/consoles"
	"go_tools/tools"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	tools.ClearConsole()
	// 持续接收控制台的输入
	// 检测用户的键盘按键，按下键盘上的esc退出
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("------------------------------")
		fmt.Printf("%s\n%s\n\n%s\n%s\n", "支持的命令(请输入编号):", "1. jsonFormat", "快捷命令:", "jof: [default]jsonFormat(将全部使用默认值,从默认文件(Documents\\json-in.json)读取并输出到默认文件(Documents\\json-out.json))")
		if tools.MonitoringKeyboard("请输入命令编号(jof):", keyboard.KeyEsc) {
			return
		}
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			if consoles.JsonFormatFromConsole("") {
				return
			}
			//  else {
			// 	consoles.ClearConsole()
			// }
		case "jof":
		default:
			input = "jof"
			if consoles.JsonFormatFromConsole(input) {
				return
			} else {
				tools.ClearConsole()
			}
		}
	}
}
