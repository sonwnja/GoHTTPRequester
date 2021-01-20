package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"math/rand"
	"os"
	"reflect"
	"time"
)
//实例化1个窗口
var window MainWindow
//分配7个LineEdit控件的指针
var urlInput *walk.LineEdit
var header1Key,header1Val *walk.LineEdit
var header2Key,header2Val *walk.LineEdit
var header3Key,header3Val *walk.LineEdit
//分配2个TextEdit控件的指针
var resp *walk.TextEdit
var head *walk.TextEdit
//分配1个ComboBox控件指针
var method *walk.ComboBox
//分配1个TextEdit控件的指针
var data *walk.TextEdit
//分配1个接口数组指针
var proxies []interface{}
//分配1个TableView控件指针
var tv *walk.TableView
//实例化1个[]FieldSetting
var FS []FieldSetting
//分配3个LineEdit控件的指针
var proxyProto *walk.LineEdit
var proxyIP *walk.LineEdit
var proxyPort *walk.LineEdit
//分配1个Webview控件指针
var render *walk.WebView
//定义1个全局的系统当前用户临时目录变量
var tmpDir string
//定义1个全局的临时文件名变量
var tmpFileName string

type METHOD struct {
	Id   int
	Method string
}

func init()  {
	tmpDir = os.Getenv("TMP")
	timestamp := time.Now().Unix()
	tmpFileName = fmt.Sprintf("%s\\.GoHTTPRequester_%d.html",tmpDir,timestamp)
	f,_ := os.Create(tmpFileName)
	err := f.Close()
	if err != nil{
		panic(err)
	}
}

func Methods() []*METHOD {
	return []*METHOD{
		{1, "GET"},
		{2, "POST"},
		{3, "HEAD"},
		{4, "PUT"},
		{5, "DELETE"},
		{6, "OPTIONS"},
	}
}

type FieldSetting struct {
	Id   int
	Proto string
	Host string
	Port string
	Speed string
}


func ProxySetting() []FieldSetting {
	return FS
}
func SetProxy(p []interface{}){
	FS = nil
	tvHDErr := tv.SetModel(nil)
	if tvHDErr != nil{
		panic(tvHDErr)
	}
	for i,v := range p{
		m := v.(map[string]interface{})
		proto := m["protocol"]
		ip := m["ip"]
		port := m["port"]
		speed := m["speed"]
		protocol := fmt.Sprintf("%s",proto)
		host := fmt.Sprintf("%s",ip)
		hostPort := fmt.Sprintf("%s",port)
		sp := fmt.Sprintf("%g",speed)
		var fs FieldSetting
		fs.Id = i + 1
		fs.Proto = protocol
		fs.Host = host
		fs.Port = hostPort
		fs.Speed = sp
		FS = append(FS, fs)
	}
}

func writeTmpHTMLFile(content string)  {
	var (
		err error
	)
	fd, err := os.OpenFile(tmpFileName, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	} else {
		_,err = fd.Write([]byte(content))
		if err != nil{
			panic(err)
		}
	}
	err = fd.Close()
	if err != nil{
		panic(err)
	}
}
func delTmpHTMLFile()  {
	err := os.Remove(tmpFileName)
	if err != nil{
		panic(err)
	}
}

func main() {
	//设置窗口标题
	window.Title = "Go HttpRequester 版本:v2.0"
	//垂直布局
	window.Layout = HBox{}
	//窗口子内容
	window.Children = []Widget{
		GroupBox{
			MaxSize: Size{Width: 500},
			Font: Font{
				Family:    "微软雅黑",
				PointSize: 0,
				Bold:      true,
			},
			Title: "请求",
			Layout: VBox{},
			Children: []Widget{
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
					},
					Title: "设置请求",
					Layout: VBox{},
					Children: []Widget{
						TabWidget{
							Pages: []TabPage{
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Title: "发请求",
									Layout: HBox{},
									Children: []Widget{
										ComboBox{
											AssignTo:      &method,
											MaxSize:       Size{Width:60},
											BindingMember: "Id",
											CurrentIndex:  0,
											DisplayMember: "Method",
											Model:         Methods(),
										},
										LineEdit{
											AssignTo: &urlInput,
											CueBanner: "输入url",
											MaxLength:          255,
										},
										PushButton{
											MaxSize: Size{Width: 50},
											Text: "请求",
											OnClicked: func() {
												method := method.Text()
												url := urlInput.Text()
												if url == ""{
													return
												}
												headers := make(map[string]string)
												h1k,h1v := header1Key.Text(),header1Val.Text()
												h2k,h2v := header2Key.Text(),header2Val.Text()
												h3k,h3v := header3Key.Text(),header3Val.Text()
												pProto,pIP,pPort := proxyProto.Text(),proxyIP.Text(),proxyPort.Text()
												var proxyFlag,h,h1,h2,h3 bool
												if pProto != "" && pIP != "" && pPort != "" {
													proxyFlag = true
												}else{
													proxyFlag = false
												}
												if h1k != "" && h1v != "" {
													headers[h1k] = h1v
													h1 = true
												}else {
													h1 = false
												}
												if h2k != "" && h2v != "" {
													headers[h2k] = h2v
													h2 = true
												}else {
													h2 = false
												}
												if h3k != "" && h3v != "" {
													headers[h3k] = h3v
													h3 = true
												}else {
													h3 = false
												}
												if h1 || h2 || h3 {
													h = true
												}else {
													h = false
												}
												switch method {
												case "GET":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														responseHeader, responseBody := ProxyGET(url,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}else {
															writeTmpHTMLFile(responseBody)
															err := render.SetURL("file:///"+tmpFileName)
															if err != nil{
																panic(err)
															}
														}
													}else {
														responseHeader, responseBody := GET(url)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}else {
															writeTmpHTMLFile(responseBody)
															err := render.SetURL("file:///"+tmpFileName)
															if err != nil{
																panic(err)
															}
														}
													}
												case "POST":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														data := data.Text()
														responseHeader, responseBody := ProxyPOST(url,data,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}else {
															writeTmpHTMLFile(responseBody)
															err := render.SetURL("file:///"+tmpFileName)
															if err != nil{
																panic(err)
															}
														}
													}else {
														data := data.Text()
														responseHeader, responseBody := POST(url,data)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}else {
															writeTmpHTMLFile(responseBody)
															err := render.SetURL("file:///"+tmpFileName)
															if err != nil{
																panic(err)
															}
														}
													}
												case "HEAD":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														data := data.Text()
														responseHeader, responseBody := ProxyHEAD(url,data,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}else {
														data := data.Text()
														responseHeader, responseBody := HEAD(url,data)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}
												case "PUT":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														data := data.Text()
														responseHeader, responseBody := ProxyPUT(url,data,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}else {
														data := data.Text()
														responseHeader, responseBody := PUT(url,data)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}
												case "DELETE":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														data := data.Text()
														responseHeader, responseBody := ProxyDELETE(url,data,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}else {
														data := data.Text()
														responseHeader, responseBody := DELETE(url,data)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}
												case "OPTIONS":
													if h{
														setHeader(headers)
													}
													if proxyFlag {
														data := data.Text()
														responseHeader, responseBody := ProxyOPTIONS(url,data,pProto,pIP,pPort)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}else {
														data := data.Text()
														responseHeader, responseBody := OPTIONS(url,data)
														headHDErr := head.SetText(responseHeader)
														if headHDErr != nil {
															panic(headHDErr)
														}
														respHDErr := resp.SetText(responseBody)
														if respHDErr != nil {
															panic(respHDErr)
														}
													}
												}
											},
										},
									},
								},
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Title: "请求头",
									Layout: VBox{},
									Children: []Widget{
										Composite{
											Layout: HBox{},
											Children: []Widget{
												LineEdit{
													AssignTo: &header1Key,
													CueBanner: "Key",
												},
												LineEdit{
													AssignTo: &header1Val,
													CueBanner: "Value",
												},
											},
										},
										Composite{
											Layout: HBox{},
											Children: []Widget{
												LineEdit{
													AssignTo: &header2Key,
													CueBanner: "Key",
												},
												LineEdit{
													AssignTo: &header2Val,
													CueBanner: "Value",
												},
											},
										},
										Composite{
											Layout: HBox{},
											Children: []Widget{
												LineEdit{
													AssignTo: &header3Key,
													CueBanner: "Key",
												},
												LineEdit{
													AssignTo: &header3Val,
													CueBanner: "Value",
												},
											},
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
						Bold:      false,
					},
					Layout: VBox{},
					Title: "设置代理",
					Children: []Widget{
						TabWidget{
							Pages: []TabPage{
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Layout: HBox{},
									Title: "填写代理",
									Children: []Widget{
										Composite{
											Layout: HBox{},
											Children: []Widget{
												LineEdit{
													AssignTo: &proxyProto,
													CueBanner: "协议",
												},
												LineEdit{
													AssignTo: &proxyIP,
													CueBanner: "IP",
													MinSize: Size{Width: 200},
												},
												LineEdit{
													AssignTo: &proxyPort,
													CueBanner: "端口",
												},
											},
										},
									},
								},
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Title: "获取代理",
									Layout: VBox{},
									Children: []Widget{
										Composite{
											Layout: HBox{},
											MaxSize: Size{Height: 30},
											Children: []Widget{
												PushButton{
													Text: "获取",
													Font: Font{PointSize: 10},
													OnClicked: func() {
														proxies = getProxy()
														SetProxy(proxies)
														tvHDErr := tv.SetModel(ProxySetting())
														if tvHDErr != nil {
															panic(tvHDErr)
														}
													},
												},
												PushButton{
													Text: "刷新",
													Font: Font{PointSize: 10},
													OnClicked: func() {
														rand.Seed(time.Now().UnixNano())
														page := rand.Intn(4)
														proxies = refresh(page+1)
														SetProxy(proxies)
														tvHDErr := tv.SetModel(ProxySetting())
														if tvHDErr != nil {
															panic(tvHDErr)
														}
													},
												},
												PushButton{
													Font: Font{PointSize: 10},
													Text: "添加",
													OnClicked: func() {
														model := tv.Model()
														index := tv.CurrentIndex()+1
														v := reflect.ValueOf(model)
														var proto,host,port string
														for i:=0;i<v.Len();i++{
															id := v.Index(i).Field(0).Int()
															if int64(index) == id{
																proto = v.Index(i).Field(1).String()
																host = v.Index(i).Field(2).String()
																port = v.Index(i).Field(3).String()
															}
														}
														proxyProHDErr := proxyProto.SetText(proto)
														if proxyProHDErr != nil {
															panic(proxyProHDErr)
														}
														proxyIPHDErr := proxyIP.SetText(host)
														if proxyIPHDErr != nil {
															panic(proxyIPHDErr)
														}
														proxyPortHDErr := proxyPort.SetText(port)
														if proxyPortHDErr != nil{
															panic(proxyPortHDErr)
														}
													},
												},
											},
										},
										TableView{
											AssignTo:           &tv,
											Columns: []TableViewColumn{
												{
													Name: "Id",
													DataMember: "Id",
													Title:      "序号",
													Width:      50,
												},
												{
													DataMember: "Proto",
													Title:      "协议",
													Width:      50,
												},
												{
													Name: "Host",
													DataMember: "Host",
													Title:      "IP",
													Width:      150,
												},
												{
													Name: "Port",
													DataMember: "Port",
													Title:      "端口",
													Width:      80,
												},
												{
													Name: "Speed",
													DataMember: "Speed",
													Title:      "响应速度",
												},
											},
											ColumnsOrderable:            true,
											ColumnsSizable:              true,
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
						Bold:      false,
					},
					Title: "POST请求数据",
					Layout: VBox{},
					Children: []Widget{
						Composite{
							Layout: VBox{},
							Children: []Widget{
								TextEdit{
									AssignTo: &data,
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
			},
			Title: "响应",
			Layout: VBox{},
			Children: []Widget{
				GroupBox{
					Font: Font{
						Family:    "微软雅黑",
						PointSize: 0,
						Bold:      false,
					},
					Title: "请求结果",
					Layout: VBox{},
					Children: []Widget{
						TabWidget{
							Pages: []TabPage{
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Layout: HBox{},
									Title: "文本",
									Children: []Widget{
										GroupBox{
											Title: "响应头",
											Layout: VBox{},
											Children: []Widget{
												TextEdit{
													AssignTo: &head,
													HScroll: true,
													VScroll: true,
													ReadOnly: true,
												},
											},
										},
										GroupBox{
											Title: "响应体",
											Layout: VBox{},
											Children: []Widget{
												TextEdit{
													AssignTo:           &resp,
													HScroll:            true,
													VScroll:            true,
													ReadOnly: true,
												},
											},
										},
									},
								},
								{
									Background: SystemColorBrush{
										Color: 4,
									},
									Title: "渲染",
									Layout: VBox{},
									Children: []Widget{
										WebView{
											AssignTo: &render,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	//运行窗口
	window.Run()
	//关闭窗口后删除临时文件
	delTmpHTMLFile()
}
