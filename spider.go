package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DUANZI struct{
	Content string
}

func HttpGet(url string) (result string,err error){
	resp,err1 := http.Get(url)
	if err1 != nil{
		err = err
		return
	}

	defer resp.Body.Close()

	buf := make([]byte,2*1024)
	for{
		n,_ := resp.Body.Read(buf)
		if n==0{
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int,page chan<- int){
	var url = "https://m.pengfue.com/index_"+strconv.Itoa(i)+".html"
	fmt.Printf("正在爬取第%d页,网址：%s",i,url)

	result,err := HttpGet(url)
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Printf("结果",result)
	re := regexp.MustCompile(`<div class="con-img">(?s:(.*?))</div>`)
	if re == nil{
		fmt.Println("正则解析错误")
		return
	}
	contents := re.FindAllStringSubmatch(result,-1)

	var content DUANZI
	var jsons []DUANZI
	for _,text := range contents{
		content = DUANZI{strings.Replace(text[1],"\n","",-1)}
		jsons = append(jsons,content)

	}

	json1,json_err := json.MarshalIndent(jsons,"","  ")
	if json_err != nil{
		fmt.Println("json格式错误")
	}
	f,f_err := os.Create("段子"+strconv.Itoa(i)+".json")
	if f_err != nil{
		fmt.Println("创建文件失败")
	}
	defer f.Close()
	var json_str = string(json1)
	f.WriteString(json_str)
	page <-i
}

func DoWork(start,end int){
	fmt.Printf("准备爬取第%d页到%d页,",start,end)
	page := make(chan int)
	for i:=start;i<=end;i++{
		go SpiderPage(i,page)
	}
	for i :=start;i<=end;i++{
		fmt.Printf("第%d个页面爬取完成",<-page)
	}
}

func main(){
	var start, end int
	fmt.Println("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入终止页(>起始页)：")
	fmt.Scan(&end)

	DoWork(start,end)
}

