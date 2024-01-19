package test

import (
	"go_tools/tools"
	"os"
	"testing"
)

type TestStruct struct {
	// 定义测试用例所需的数据
}

func (ts *TestStruct) TestMethod(t *testing.T) {
	// 安排测试用例
	// 调用待测试的方法
	// 读取用户home路径
	userHome := os.Getenv("USERPROFILE")
	var basePath string = userHome + "\\Documents"
	var inFilePath string = basePath + "\\json-in.json"
	var outFilePath string = basePath + "\\json-out.json"
	tools.JsonFormatFromFile(inFilePath, outFilePath)
	// 断言测试结果
}
func Test(t *testing.T) {
	// 创建一个测试结构
	ts := &TestStruct{}
	// 运行测试方法
	ts.TestMethod(t)
}
