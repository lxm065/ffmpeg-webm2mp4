package main

import (
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建新的应用
	myApp := app.New()

	// 创建窗口
	myWindow := myApp.NewWindow("FFmpeg Converter")

	// 窗口大小
	myWindow.Resize(fyne.NewSize(1200, 600))
	// 窗口居中
	myWindow.CenterOnScreen()
	// 窗口标题
	myWindow.SetTitle("FFmpeg 转换器")

	selectButton_text := widget.NewEntry()
	selectButton_text.SetPlaceHolder("you select file name will be show here,eg:input.webm")

	// 输入输出文件名称的输入框
	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("input output file name,eg:output.mp4")

	startButton_text := widget.NewEntry()
	startButton_text.SetPlaceHolder("")

	// 运行按钮
	runButton := widget.NewButton("run now", func() {
		// 提示开始运行
		startButton_text.SetText("start running,please wait...")

		// 判断输入文件是否为空是否是.webm文件
		if selectedFile == "" || !strings.HasSuffix(selectedFile, ".webm") {
			dialog.ShowInformation("error", "please select input file or it's a .webm file", myWindow)
			return
		}
		// 判断输出文件是否为空是否是.mp4文件
		if outputEntry.Text == "" || !strings.HasSuffix(outputEntry.Text, ".mp4") {
			dialog.ShowInformation("error", "please input output file name or it's a .mp4 file", myWindow)
			return
		}

		// 获取输入文件路径
		inputFile := outputEntry.Text
		// 构建并执行FFmpeg命令
		cmd := exec.Command("ffmpeg", "-i", selectedFile, inputFile)
		err := cmd.Run()
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("success", "translate complete", myWindow)
		// 提示运行结束
		startButton_text.SetText("run over")
	})

	// 文件选择器
	fileDialog := dialog.NewFileOpen(
		func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if r == nil {
				return
			}
			selectedFile = r.URI().Path()
			fmt.Println("select file:", selectedFile)
			// 文件选择器路径显示在文本框中
			selectButton_text.SetText(selectedFile)
		}, myWindow)

	// 文件选择器过滤器
	selectButton := widget.NewButton("select file", func() {
		fileDialog.Show()
	})

	// 将控件添加到窗口
	myWindow.SetContent(container.NewVBox(
		selectButton,
		selectButton_text,
		outputEntry,
		runButton,
		startButton_text,
	))

	myWindow.ShowAndRun()
}

// 用于存储选择的文件路径
var selectedFile string
