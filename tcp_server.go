package main

/**
tcp服务端
*/
import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

//处理用户请求
func handleConn(conn net.Conn) {
	//处理完毕关闭
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect sucessful")

	buf := make([]byte, 1024)

	for {
		//读取客户端发送内容
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("request err = ", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		fmt.Println("len = ", len(string(buf[:n])))
		//判断客户端发送内容为exit时端口客户端连接
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, " exit")
			return
		}
		//返回输入的内容，转字符串并小写转大写
		r := make(map[string]interface{})
		r["code"] = 200
		r["message"] = ""
		r["data"] = map[string]string{"content": strings.ToUpper(string(buf[:n]))}
		j, _ := json.MarshalIndent(r, "", " ")
		conn.Write([]byte(j))
	}

}

func main() {
	//服务端监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
			return
		}
		//处理用户请求，新建一个协程
		go handleConn(conn)
	}
}
