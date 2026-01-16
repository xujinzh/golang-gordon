package main

import (
    "os"
    "github.com/gotk3/gotk3/gtk"
)

func main() {
    gtk.Init(&os.Args)

    builder, err := gtk.BuilderNewFromFile("color.glade")
    if err != nil {
        panic(err)
    }

    // 在 gotk3 中，直接获取对象并转换为窗口
    obj, err := builder.GetObject("window1")
    if err != nil {
        panic(err)
    }
    win := obj.(*gtk.Window) // 进行类型断言

    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    win.ShowAll()
    gtk.Main()
}