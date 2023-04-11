package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const DistPath = "D:\\2345Downloads\\znb921\\"
const StartPage = "https://www.169w.cc/c49.aspx"
const BasePro = "https://www.169w.cc/"

var wg sync.WaitGroup

type Page struct {
	title string
	url   string
}

type Girl struct {
	title string
	urls  []string
}

func RandStr2() string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 3)
	rand.Read(result)
	return hex.EncodeToString(result)
}

func downloadPic(url string, count string, cate string) {
	s := DistPath + cate + "" +
		"//" + RandStr2() + ".jpg"
	if !isExist(DistPath + cate + "//") {
		_ = os.MkdirAll(DistPath+cate+"//", os.ModePerm)
	}
	url = BasePro + url
	rsp, err := http.Get(url)
	reader := bufio.NewReader(rsp.Body)
	if rsp.StatusCode != 200 {
		fmt.Println(url + " code --> " + strconv.Itoa(rsp.StatusCode))
		return
	}
	file, err := os.OpenFile(s, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = reader.WriteTo(file)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(url+" --> ok -->", count)
	}
}

func ParsePage(page string) []Page {
	client := &http.Client{}
	var resList []Page
	req, err := http.NewRequest("GET", page, nil) //建立连接
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive") //设置请求头
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req) //拿到返回的内容
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("fatal err")
		log.Fatal(err)
	}
	docDetail.Find(".nl li"). //定位到html页面指定元素
					Each(func(i int, s *goquery.Selection) { //循环遍历每一个指定元素
			src, _ := s.Find("a").Attr("href")
			title, _ := s.Find("a").Attr("title")
			resList = append(resList, Page{
				url:   src,
				title: title,
			})
		})
	return resList
}

func parseHtml(url string) []string {
	var urls []string
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil) //建立连接
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive") //设置请求头
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req) //拿到返回的内容
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("fatal err")
		log.Fatal(err)
	}
	docDetail.Find("#content div div img"). //定位到html页面指定元素
						Each(func(i int, s *goquery.Selection) { //循环遍历每一个指定元素
			src, _ := s.Attr("src")
			urls = append(urls, src)
		})
	return urls
}

func parseThree(url string, title string) Girl {
	url = BasePro + url
	urls1 := parseHtml(url)
	url = strings.Replace(url, ".aspx", "", 1) + "p2" + ".aspx"
	urls2 := parseHtml(url)
	url = strings.Replace(url, ".aspx", "", 1) + "p3" + ".aspx"
	urls3 := parseHtml(url)
	urls1 = append(urls1, urls2...)
	urls1 = append(urls1, urls3...)
	wg.Done()
	return Girl{
		title: title,
		urls:  urls1,
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func main() {
	var allGirl []Girl
	pages := ParsePage(StartPage)
	if len(pages) == 0 {
		fmt.Println("00000000")
		return
	}
	fmt.Println(len(pages), "pages")
	wg.Add(len(pages))
	for i := 0; i < len(pages); i++ {
		go func(i int) {
			aGirl := parseThree(pages[i].url, pages[i].title)
			fmt.Println(len(aGirl.urls), "a girl len, index -->", i)
			allGirl = append(allGirl, aGirl)
		}(i)
	}
	wg.Wait()
	fmt.Println(len(allGirl), " let's go")
	for i := 0; i < len(allGirl); i++ {
		for j := 0; j < len(allGirl[i].urls); j++ {
			downloadPic(allGirl[i].urls[j], strconv.Itoa(i)+" / "+strconv.Itoa(j), allGirl[i].title)
			time.Sleep(time.Second / 10)
		}
	}
}
