package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Species struct {
	Id   int
	Method string
}

func KnownSpecies() []*Species {
	return []*Species{
		{1, "GET"},
		{2, "POST"},
		{3, "HEAD"},
		{4, "PUT"},
		{5, "DELETE"},
		{6, "OPTIONS"},
	}
}

func main() {
	//实例化一个窗口
	var window MainWindow
	//分配一个LineEdit控件的指针
	var le *walk.LineEdit
	//设置窗口标题
	window.Title = "Go HttpRequester"
	//垂直布局
	window.Layout = HBox{}
	//窗口子内容
	window.Children = []Widget{
		GroupBox{
			Font: Font{
				Family:    "微软雅黑",
				PointSize: 0,
				Bold:      true,
				Italic:    false,
				Underline: false,
				StrikeOut: false,
			},
			Title: "请求",
			Layout: VBox{},
			Children: []Widget{
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
						Italic:    false,
						Underline: false,
						StrikeOut: false,
					},
					Title: "设置请求URL",
					Layout: VBox{},
					Children: []Widget{
						Composite{
							Layout: HBox{},
							Children: []Widget{
								ComboBox{
									BindingMember:         "Id",
									CurrentIndex:          0,
									DisplayMember:         "Method",
									Model:                 KnownSpecies(),
								},
								Label{
									Text: "URL:",
								},
								LineEdit{
									AssignTo: &le,
									ToolTipText:        "输入url",
									MaxLength:          255,
								},
								PushButton{
									Text: "请求",
									OnClicked: func() {
										fmt.Println(le.Text())
									},
								},
							},
						},
					},
				},
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
						Italic:    false,
						Underline: false,
						StrikeOut: false,
					},
					Title: "设置请求头",
					Layout: VBox{},
					Children: []Widget{
						Composite{
							Layout: HBox{},
							Children: []Widget{
								Label{Text: "Key"},
								LineEdit{},
								Label{Text: "Value"},
								LineEdit{},
							},
						},
						Composite{
							Layout: HBox{},
							Children: []Widget{
								Label{Text: "Key"},
								LineEdit{},
								Label{Text: "Value"},
								LineEdit{},
							},
						},
						Composite{
							Layout: HBox{},
							Children: []Widget{
								Label{Text: "Key"},
								LineEdit{},
								Label{Text: "Value"},
								LineEdit{},
							},
						},
					},
				},
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
						Italic:    false,
						Underline: false,
						StrikeOut: false,
					},
					Title: "POST请求数据",
					Layout: VBox{},
					Children: []Widget{
						Composite{
							Layout: VBox{},
							Children: []Widget{
								TextEdit{
								},
							},
						},
					},
				},
			},
		},
		GroupBox{
			Font: Font{
				Family:    "微软雅黑",
				PointSize: 0,
				Bold:      true,
				Italic:    false,
				Underline: false,
				StrikeOut: false,
			},
			Title: "响应",
			Layout: VBox{},
			Children: []Widget{
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
						Italic:    false,
						Underline: false,
						StrikeOut: false,
					},
					Title: "请求结果",
					Layout: VBox{},
					Children: []Widget{
						GroupBox{
							Title: "响应头",
							Layout: VBox{},
							Children: []Widget{
								TextEdit{},
							},
						},
						GroupBox{
							Title: "响应体",
							Layout: VBox{},
							Children: []Widget{
								TextEdit{},
							},
						},
					},
				},
			},
		},
	}
	//运行
	window.Run()
}
