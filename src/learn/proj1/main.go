package main
import (
    "encoding/gob"
    "fmt"
    "os"
)
func main() {
    info := "how are you?"
    file, err := os.Create("./data/btext.gob")
    if err != nil {
        fmt.Println("文件创建失败", err.Error())
        return
    }
    defer file.Close()
    encoder := gob.NewEncoder(file)
    err = encoder.Encode(info)
    if err != nil {
        fmt.Println("编码错误", err.Error())
        return
    } else {
        fmt.Println("编码成功")
    }

	file1, err1 := os.Open("./data/btext.gob")
    if err1 != nil {
        fmt.Println("文件打开失败", err1.Error())
        return
    }
    defer file1.Close()
    decoder := gob.NewDecoder(file1)
    info1 := ""
    err1 = decoder.Decode(&info1)
    if err1 != nil {
        fmt.Println("解码失败", err1.Error())
    } else {
        fmt.Println("解码成功")
        fmt.Println(info)
    }
}