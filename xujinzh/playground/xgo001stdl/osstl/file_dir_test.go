package osstl_test

import (
	"fmt"
	"os"
	"testing"
)

// panic 介绍
// 主动抛出不可恢复的严重错误，导致程序中断执行。仅用于不可预测的重大错误（如系统运行状态不一致、严重配置错误）。普通业务错误应优先使用error
// 配合defer, recover可以在程序崩溃前进行必要的清理工作。
// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("恢复了:", r) // 捕获并处理panic
// 		}
// 	}()

// 	panic("发生严重错误") // 主动触发
// 	fmt.Println("这行代码不会执行")
// }

// 测试创建一个文件
func TestCreate(t *testing.T) {
	// 创建一个文件。内部就是调用：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// 1. 文件不存在会创建一个新文件，O_CREATE
	// 2. 文件存在会先清空文件再创建新文件，O_TRUNC
	// 特别注意，当待创建的文件路径名的上一层目录不存在，那么创建失败，报 no such file or directory
	f, err := os.Create("./assets/test.txt")
	if err != nil { // 如果创建失败，报异常
		// 触发panic
		panic(err)
	}
	// 记得关闭文件
	defer f.Close()
	// todo
	defer fmt.Printf("程序运行结束。")
}

// 测试创建单级目录
func TestMkdir(t *testing.T) {
	// 创建一个文件夹，而不是多级目录。注意非递归创建，即创建的文件夹上一层目录必须存在
	err := os.Mkdir("./assets/test", os.ModePerm)
	// 如果创建失败，那么报异常信息
	if err != nil {
		panic(err)
	}
}

// 测试创建多级目录
func TestMkdirAll(t *testing.T) {
	// 递归创建多级目录
	err := os.MkdirAll("./assets/a/b", os.ModePerm)
	// 如果创建失败，那么报异常信息
	if err != nil {
		panic(err)
	}
}

// 测试删除文件或空目录
func TestRemove(t *testing.T) {
	// 删除空目录
	err := os.Remove("./assets/test")
	// 如果删除空目录失败
	if err != nil {
		t.Log("删除空目录失败：", err)
	}
	// 删除文件
	err = os.Remove("./assets/test.txt")
	// 如果删除文件失败
	if err != nil {
		t.Log("删除文件失败：", err)
	}

}

// 测试删除多级目录
func TestRemoveAll(t *testing.T) {
	// 删除多级目录
	err := os.RemoveAll("./assets/a")
	// 如果删除多级目录失败
	if err != nil {
		t.Log("删除目录失败：", err)
	}
}
