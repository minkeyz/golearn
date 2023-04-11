package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

//https://backend.uthash.com/register
//?account_id=susu101&address=TBG6F9nfTmZjoi8AmdovZptuC59a135AQ1
//&password=e00fb22188e2e3461b94108791da4b36&verify_password=e00fb22188e2e3461b94108791da4b36&
//invite_code=&url=uthash.com

type Result struct {
	Data Token `json:"data"`
}

type Token struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}

func generateKeyPair() (b5 string, pk string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	address = "41" + address[2:]
	addb, _ := hex.DecodeString(address)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}
	return base58.Encode(addb), hexutil.Encode(privateKeyBytes)[2:]
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func genString(flag bool) string {
	rand.Seed(time.Now().UnixNano())
	str := strings.Builder{}
	length := len(CHARS)
	ll := 6
	if flag {
		time.Sleep(1e4)
		ll = 5
		str.WriteString(CHARS[rand.Intn(25)])
	}
	for i := 0; i < ll; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

func genUrl() string {
	ad, _ := generateKeyPair()
	ps := Md5("12312345")
	rand.Seed(time.Now().UnixNano())
	un := fmt.Sprintf("%d", rand.Int63())
	baseUrl := "https://backend.uthash.com/register"
	acString := fmt.Sprintf("?account_id=%s", un)
	adString := fmt.Sprintf("&address=%s", ad)
	psString := fmt.Sprintf("&password=%s", ps)
	psString2 := fmt.Sprintf("&verify_password=%s", ps)
	inviteString := "&invite_code=&url=uthash.com"
	return baseUrl + acString + adString + psString + psString2 + inviteString
}

func requestZ(url string) (code int) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseCode := response.StatusCode
	//buf := bytes.NewBuffer(make([]byte, 0, 128))
	//buf.ReadFrom(response.Body)
	//responseString := string(buf.Bytes())
	return responseCode
}

func manager() {
	reqUrl := genUrl()
	fmt.Printf("url --> %s \n", reqUrl)
	requestZ(reqUrl)
}

func Multipart(count int) {
	url := "https://api.6689.vip/api/auth/register"
	rn := genString(true)
	ps := genString(false)
	client := &http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//添加表单属性
	// 参数1 （File参数）
	//添加您的镜像文件
	// 参数2 fileType (普通参数)
	fileWriter1, _ := bodyWriter.CreateFormField("name")
	_, _ = fileWriter1.Write([]byte(rn))

	fileWriter2, _ := bodyWriter.CreateFormField("password")
	_, _ = fileWriter2.Write([]byte(ps))

	fileWriter3, _ := bodyWriter.CreateFormField("password_confirmation")
	_, _ = fileWriter3.Write([]byte(ps))

	_ = bodyWriter.Close()
	req, _ := http.NewRequest("POST", url, bodyBuf)
	//添加头文件
	//req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	//获取返回值
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	//defer func(Body io.ReadCloser) {
	//	_ = Body.Close()
	//}(resp.Body)
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, count, rn, ps, string(respBytes))
	if resp.StatusCode != 200 {
		fmt.Println("code: --->", resp.StatusCode)
		time.Sleep(1e10)
	}
	Login(rn, ps)
}

func Login(name, password string) {
	url := "https://api.6689.vip/api/auth/login"
	client := &http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	field1, _ := bodyWriter.CreateFormField("name")
	_, _ = field1.Write([]byte(name))

	field2, _ := bodyWriter.CreateFormField("password")
	_, _ = field2.Write([]byte(password))

	defer bodyWriter.Close()
	req, _ := http.NewRequest("POST", url, bodyBuf)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("login Ok: ", name)
	var r Result
	_ = json.Unmarshal(respBytes, &r)
	Bound(r.Data.AccessToken, r.Data.TokenType)

}

func Bound(token, bear string) {
	url := "https://api.6689.vip/api/auth/update_wallet"
	ad, _ := generateKeyPair()
	client := &http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	field1, _ := bodyWriter.CreateFormField("wallet_address")
	_, _ = field1.Write([]byte(ad))

	defer bodyWriter.Close()
	req, _ := http.NewRequest("POST", url, bodyBuf)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("Authorization", bear+" "+token)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("Bound Ok: ", ad)
	fmt.Println("---------------------------------------------------------")
}

func SamplePost() {
	rand.Seed(time.Now().UnixNano())
	info := make(map[string]string)
	info["name"] = "zx7888"
	info["password"] = "123456"
	info["password_confirmation"] = "123456"
	info["funds_password"] = "222"
	info["wallet_address"] = "111"
	fmt.Println(info)

	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bytes.NewReader(bytesData)

	url := "https://api.6689.vip/api/auth/register" //要访问的Url地址
	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryGYsN7uoJYEX1wN8N")
	request.Header.Set("Origin", "https://98hash.vip/")
	request.Header.Set("Referer", "https://98hash.vip/")
	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{}
	resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println(err)
		return
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, string(respBytes))
}

func main() {
	//for i := 0; i < 120000; i++ {
	//	Multipart(i)
	//}
	url := "https://www.toutiaoyule.com/m/ajaxlistredis.php?from=tt07&cate=83&pagesize=20&page=0&id=37098659"
	var count = 0
	for i := 0; i < 12; i++ {
		go func() {
			for {
				code := requestZ(url)
				count++
				fmt.Println("count -->", count, "---->", code)
			}
		}()
	}

	time.Sleep(time.Hour)
}
