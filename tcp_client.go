package main

/**
tcp 客户端
*/
import (
	"fmt"
	"net"
	"os"
)

func main() {
	//连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("conn err = ", err)
		return
	}
	//调用完毕，关闭
	defer conn.Close()

	go func() {
		//从键盘输入内容，给服务器发送
		str := make([]byte, 1024)
		fmt.Println("请输入内容：")
		for {

			n, err := os.Stdin.Read(str) //读取
			if err != nil {
				fmt.Println("request err = ", err)
				return
			}
			//发送输入的内容
			conn.Write(str[:n])
		}
	}()

	//接收服务端返回内容
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务端请求
		if err != nil {
			fmt.Println("response err = ", err)
			return
		}
		fmt.Println("服务端输出：" + string(buf[:n])) //打印接收到的内容
	}
}
