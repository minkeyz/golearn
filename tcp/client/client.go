package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdef0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:23456")
	if err != nil {
		panic("dial error: " + err.Error())
	}
	reader := bufio.NewReader(os.Stdin)
	rans := getRandstring(6)
	for {
		fmt.Print("请输入：")
		msg, _ := reader.ReadString('\n')
		msg = rans + strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			panic("write error")
		}
	}
	conn.Close()
}
