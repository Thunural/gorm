package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gormui/config"
	"gormui/gen"
	"gormui/theme"
	"gormui/utils"
	"os"
	"strings"
)

func main() {
	var p config.Param
	myApp := app.New()
	myApp.Settings().SetTheme(&theme.MyTheme{})
	myWindow := myApp.NewWindow("Gorm代码生成")
	myWindow.Resize(fyne.NewSize(440, 320))

	Address := widget.NewEntry()
	Address.SetPlaceHolder("请输入数据库地址")
	Address.Validator = validation.NewRegexp("^((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))$", "请输入正确的IP地址")

	Port := widget.NewEntry()
	Port.SetPlaceHolder("请输入数据库端口")
	Port.Validator = validation.NewRegexp("^[0-9]", "请输入正确的端口")

	Database := widget.NewEntry()
	Database.SetPlaceHolder("请输入要生成的数据库名")
	Database.Validator = validation.NewRegexp("^[a-zA-Z0-9_.]+$", "请输入正确的数据库名")

	UserName := widget.NewEntry()
	UserName.SetPlaceHolder("请输入数据库用户名")
	UserName.Validator = validation.NewRegexp("^[a-zA-Z0-9_]+$", "请输入用户名，可包含下划线")

	Password := widget.NewPasswordEntry()
	Password.SetPlaceHolder("请输入密码")
	Password.Validator = validation.NewRegexp("^[a-zA-Z0-9_~!@#$%^&*()_\\-+=<>?:\"{}|,.\\/;'\\\\[\\]·~！@#￥%……&*（）——\\-+={}|《》？：“”【】、；‘'，。、]{1,20}", "请输入数据库密码")

	ProjectName := widget.NewEntry()
	ProjectName.SetPlaceHolder("请输入包名")
	ProjectName.Validator = validation.NewRegexp("^[a-zA-Z]+$", "请输入包名，只允许输入英文")

	genType := widget.NewCheckGroup([]string{"model", "dao", "service"}, func(i []string) {
		p.CheckType = i
	})

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "数据库地址", Widget: Address},
			{Text: "数据库端口", Widget: Port},
			{Text: "数据库名", Widget: Database},
			{Text: "用户名", Widget: UserName},
			{Text: "密码", Widget: Password},
			{Text: "包名", Widget: ProjectName},
			{Text: "类型", Widget: genType},
		},

		OnSubmit: func() { // optional, handle form submission
			// 赋值给到结构体方便使用
			p.Address = Address.Text
			p.Port = Port.Text
			p.Database = Database.Text
			p.UserName = UserName.Text
			p.Password = Password.Text
			p.ProjectName = ProjectName.Text
			if len(p.CheckType) < 1 || p.OutFiles == "" {
				dialog.ShowInformation("错误", "请选择生成类型或输出路径", myWindow)
				return
			}
			// 创建一个提示框
			confirm := dialog.NewConfirm("提示", "是否生成项目？", func(check bool) {
				if check {
					err := gen.Check(p)
					// 如果返回等于nil则目录存在
					if err == nil {
						path := p.OutFiles + "/" + p.ProjectName
						tips := dialog.NewConfirm("提示", path+"\n目录已存在是否重新生成？", func(check bool) {
							if check {
								_ = os.RemoveAll(path)
								err := gen.Project(p)
								if err != nil {
									dialog.ShowInformation("错误", err.Error(), myWindow)
									return
								}
								dialog.ShowInformation("完成", "已重新生成完毕", myWindow)
							}
						}, myWindow)
						tips.Show()
					} else {
						err := gen.Project(p)
						if err != nil {
							dialog.ShowInformation("错误", err.Error(), myWindow)
							return
						}
						dialog.ShowInformation("完成", "已生成完毕", myWindow)
					}
				}
			}, myWindow)
			// 将提示框显示出来
			confirm.Show()
		},
		SubmitText: "提交",
	}
	form.Append("输出路径", widget.NewButton("浏览", func() {
		fileBrowse := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			filePath := utils.String(uri)
			p.OutFiles = strings.Replace(filePath, "file://", "", -1)
		}, myWindow)
		fileBrowse.Show()
	}))
	lay := layout.NewPaddedLayout()
	centered := container.New(lay, form)
	myWindow.SetContent(centered)
	myWindow.ShowAndRun()
}
