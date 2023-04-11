package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
	"unicode"
)

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}

func writer(s string) {
	f1, err1 := os.OpenFile("./res.txt", os.O_APPEND, 0666)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	defer func(f1 *os.File) {
		err := f1.Close()
		if err != nil {
		}
	}(f1)
	n, err1 := io.WriteString(f1, s+"\r\n")
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	fmt.Printf("写入 %d 个字节\n", n)
}

//func generateKeyPair() (b5 string, pk string) {
//	privateKey, _ := crypto.GenerateKey()
//	privateKeyBytes := crypto.FromECDSA(privateKey)
//	publicKey := privateKey.Public()
//	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
//	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
//	address = "41" + address[2:]
//	addb, _ := hex.DecodeString(address)
//	hash1 := s256(s256(addb))
//	secret := hash1[:4]
//	for _, v := range secret {
//		addb = append(addb, v)
//	}
//	return base58.Encode(addb), hexutil.Encode(privateKeyBytes)[2:]
//}

func generateKeyPair() (ad string, pk string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()[2:]
	return address, hexutil.Encode(privateKeyBytes)[2:]
}

/*
strings.HasSuffix(ad, "999999999") ||
		strings.HasSuffix(ad, "888888888") ||
		strings.HasSuffix(ad, "666666666") ||
		strings.HasSuffix(ad, "777777777") ||
*/
func manager(count *int64) {
	ad, pk := generateKeyPair()
	writer(ad + "----" + pk)
	fmt.Println("get one!")
	if strings.HasSuffix(ad, "19888888") ||
		strings.HasSuffix(ad, "20201130") ||
		strings.HasSuffix(ad, "12345678") {
		writer(ad + "----" + pk)
		fmt.Println("get one!")
	}
	*count++
}

func printer(count *int64) {
	time.Sleep(10 * time.Second)
	fmt.Println("count -->", *count)
	printer(count)
}

func isDigit(str string) bool {
	ss := str[2:]
	for _, x := range []rune(ss) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func main() {
	var count int64 = 0
	runtime.GOMAXPROCS(4)
	for i := 0; i < 4; i++ {
		go func(c *int64) {
			for {
				manager(c)
			}
		}(&count)
	}
	printer(&count)
	time.Sleep(720 * time.Hour)
}
