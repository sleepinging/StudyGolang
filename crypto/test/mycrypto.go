package main

import (
	"crypto/des"
	"crypto/cipher"
	"bytes"
	"encoding/base64"
	"fmt"
)

var key = []byte("19961102")
var iv = []byte("19960728")

var block cipher.Block

func init() {
	var err error
	block, err = des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		//return nil, err
	}
}

func DESEncrypt(data, key []byte, iv []byte) ([]byte, error) {
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	return out, nil
}
func DESDecrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	blockMode := cipher.NewCBCDecrypter(block, iv)
	out := make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	out = PKCS5UnPadding(out)
	return out, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func MyEncrypt(str string)(res string){
	out, _ := DESEncrypt([]byte(str), key, iv)
	res = base64Encode(out)
	return
}

func MyDecrypt(str string)(res string){
	enbyte, _ := base64Decode(str)
	//log.Println("base:", enbyte)
	out, _ := DESDecrypt(enbyte, key, iv)
	res=string(out)
	return
}

func testenc(str string){
	nstr:=MyEncrypt(str)
	fmt.Println(nstr)
	fmt.Println(MyDecrypt(nstr))
}

func main() {
	testenc("hello whh! do you have lunch?")
}
