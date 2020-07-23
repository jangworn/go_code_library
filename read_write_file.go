package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(filename, content string) {
	_, err := os.Stat(filename)
	var f *os.File
	var err1 error
	if err == nil {
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0660)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
	} else {
		f, err1 = os.Create(filename)
	}
	defer f.Close()

	_, err = f.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err)
	}

}

func readFile(filename string) (result string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//一次性读取
	/* buf := make([]byte, 1024*1)
	n, err := f.Read(buf) */

	//逐行读取
	read := bufio.NewReader(f)

	for {
		buf, err := read.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		result += string(buf)
	}

	return
}

func main() {
	var i = 0
	var content string
	for {
		if i > 10 {
			break
		}
		fmt.Println("请输入要记录的内容：")
		fmt.Scan(&content)
		writeFile("./demo.txt", content)
		i++
	}
	result := readFile("./demo.txt")
	fmt.Println(result)
}
