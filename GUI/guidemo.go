package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		// 生成：文本框
		name := ui.NewEntry()
		// 生成：标签
		greeting := ui.NewLabel(``)
		// 生成：按钮
		button := ui.NewButton(`欢迎`)
		// 设置：按钮点击事件
		button.OnClicked(func(*ui.Button) {
			greeting.SetText(`你好，` + name.Text() + `！`)
		})
		// 生成：垂直容器
		box := ui.NewVerticalBox()

		// 往 垂直容器 中添加 控件
		box.Append(ui.NewLabel(`请输入你的名字：`), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)

		// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
		window := ui.NewWindow(`你好`, 200, 100, true)

		// 窗口容器绑定
		window.SetChild(box)

		// 设置：窗口关闭时
		window.OnClosing(func(*ui.Window) bool {
			// 窗体关闭
			ui.Quit()
			return true
		})

		// 窗体显示
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}



