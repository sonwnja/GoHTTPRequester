package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)
var heads = make(map[string]string)
func setHeader(h map[string]string)  {
	heads = h
}
func GET(url string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func ProxyGET(urlLink string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("GET", urlLink,nil)
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func POST(url string,params string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func ProxyPOST(urlLink string,params string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("POST", urlLink,strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}


func HEAD(url string,params string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("HEAD", url, strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func ProxyHEAD(urlLink string,params string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("HEAD", urlLink,strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	req.Header.Add("User-Agent", "Go Http Requester")
	heads = nil
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}


func PUT(url string,params string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}


func ProxyPUT(urlLink string,params string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("PUT", urlLink,strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func DELETE(url string,params string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func ProxyDELETE(urlLink string,params string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("DELETE", urlLink,strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}



func OPTIONS(url string,params string)  (string,string){
	client := &http.Client{}
	req, _ := http.NewRequest("OPTIONS", url, strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}

func ProxyOPTIONS(urlLink string,params string,pProto string,pHost string,pPort string)  (string,string){
	urli := url.URL{}
	proxyUrl := fmt.Sprintf("%s://%s:%s",pProto,pHost,pPort)
	proxy, _ := urli.Parse(proxyUrl)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ := http.NewRequest("OPTIONS", urlLink,strings.NewReader(string(params)))
	//增加header选项
	if heads != nil{
		for k,v := range heads{
			req.Header.Add(k, v)
		}
	}
	heads = nil
	req.Header.Add("User-Agent", "Go Http Requester")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		now := time.Now().Format("2006-01-02 15:04:05")
		now = fmt.Sprintf("【错误日志】\r\n[%s] ",now)
		return "",now + err.Error()
	}
	defer response.Body.Close()
	codeDesc := response.Status
	body, _ := ioutil.ReadAll(response.Body)
	head := response.Header
	h := ""
	h += fmt.Sprintf("HTTP/1.1 %s\r\n",codeDesc)
	for k,v := range head {
		value := ""
		for i,val := range v{
			value += val
			if len(v)  > 1 && i != (len(v) -1) {
				value += ","
			}
		}
		h += fmt.Sprintf("%s: %s\r\n",k,value)
	}
	return h,string(body)
}
