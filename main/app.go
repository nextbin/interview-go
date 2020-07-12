package main

import "github.com/nextbin/hello-leetcode-go/leetcode/medium"


import (
	"strconv"
	"encoding/json"
	"fmt"
)

func main2() {
	medium.Run0098()
}


func main3() {
	c := make(chan int, 1)//修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}


type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main4() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}