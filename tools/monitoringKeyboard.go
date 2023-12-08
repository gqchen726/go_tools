package tools

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

// MonitoringKeyboard monitors the keyboard input and returns true if the ESC key is pressed, and false otherwise.
//
// It does not take any parameters.
// It returns a boolean value.
func MonitoringKeyboard(continueStr string, distKey keyboard.Key) bool {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n-->请按下ESC(退出)/ENTER(继续)<--")
	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}
	if key == distKey {
		fmt.Println("再次按下ESC确认退出")
		defer keyboard.Close()
		ClearConsole()
		return true
	} else {
		fmt.Println(continueStr)
		defer keyboard.Close()
		return false
	}
}

func ReturnMonitoringKeyboard() keyboard.Key {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n-->请按下ESC(退出)/BACKSPACE(返回)/ENTER(继续)<--")
	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}
	return key
}
