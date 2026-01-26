package osstl

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
	// 最后一个defer，在函数运行完后先执行
	defer fmt.Printf("程序运行结束。\n")
}

// 测试创建单级目录
func TestMkdir(t *testing.T) {
	// 创建一个文件夹，而不是多级目录。注意非递归创建，即创建的文件夹上一层目录必须存在
	err := os.Mkdir("./assets", os.ModePerm)
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

// 测试删除多级目录和目录中的文件
func TestRemoveAll(t *testing.T) {
	// 删除多级目录和目录中的文件
	err := os.RemoveAll("./assets/a")
	// 如果删除多级目录和目录中的文件失败
	if err != nil {
		t.Log("删除目录失败：", err)
	}
}

// 测试获取当前路径
func TestGetwd(t *testing.T) {
	// 获取当前路径
	wd, err := os.Getwd()
	// 如果失败，打印错误信息
	if err != nil {
		t.Log("获取当前路径失败：", err)
	}
	// 打印当前路径
	fmt.Printf("当前工作目录: %v\n", wd)

}

// 测试切换目录
func TestChdir(t *testing.T) {
	// 尝试切换到指定目录
	err := os.Chdir("/opt")
	// 如果切换失败
	if err != nil {
		t.Log("切换目录失败：", err)
	} else { // 切换成功
		t.Log("切换目录成功。")
	}
	// 获取切换后的工作目录
	wd, err := os.Getwd()
	// 如果获取工作目录失败
	if err != nil {
		t.Log("获取当前目录失败：", err)
	}
	// 打印当前目录
	fmt.Printf("wd: %v\n", wd)
	// fmt.Println(os.Getwd())
}

// 测试获取临时目录
func TestTempDir(t *testing.T) {
	// 获取存放临时文件的目录，不同操作系统结果可能不同
	tmp := os.TempDir()
	fmt.Printf("tmp: %v\n", tmp)
}

// 测试重命名文件
func TestRename(t *testing.T) {
	// // 重命名文件
	// err := os.Rename("./assets/test.txt", "./assets/text.txt")
	// // 捕获重命名失败的异常
	// if err != nil {
	// 	t.Log("重命名文件失败：", err)
	// }
	// 重命名文件夹
	err := os.Rename("./assets", "./resources")
	if err != nil { // 如果失败
		t.Log("重命名文件夹失败。")
	} else { // 成功
		t.Log("重命名文件夹成功。")
	}
}

// 修改文件或文件夹权限
func TestChmod(t *testing.T) {
	// 尝试改变文件的权限
	err := os.Chmod("./resources/text.txt", 0o400)
	if err != nil { // 如果失败
		t.Log("改变文件权限失败:", err)
	} else { // 成功
		t.Log("改变文件权限成功。")
	}
}
