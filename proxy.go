package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlApi = "https://ip.jiangxianli.com/api/proxy_ips"

func getProxy() []interface{}  {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlApi, nil)
	//增加header选项
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var data interface{}
	json.Unmarshal(body,&data)
	m := data.(map[string]interface{})
	data = m["data"]
	m = data.(map[string]interface{})
	data = m["data"]
	//取出数组
	list := data.([]interface{})
	return list
}

func refresh(page int) []interface{} {
	client := &http.Client{}
	refreshUrl := fmt.Sprintf( "%s?page=%d",urlApi,page)
	req, _ := http.NewRequest("GET", refreshUrl, nil)
	//增加header选项
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var data interface{}
	json.Unmarshal(body,&data)
	m := data.(map[string]interface{})
	data = m["data"]
	m = data.(map[string]interface{})
	data = m["data"]
	//取出数组
	list := data.([]interface{})
	return list
}
