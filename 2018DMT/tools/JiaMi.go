package tools

import (
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

//加密字符串
func Encrypt(strMesg, key string) (string, error) {
	//key := "1234567891234567"
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	uEnc := base64.URLEncoding.EncodeToString(encrypted)
	return string(uEnc), nil
}

//解密字符串
func Decrypt(src, key string) (strDesc string, err error) {
	uDec, _ := base64.URLEncoding.DecodeString(src)
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(uDec))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, uDec)
	return string(decrypted), err
}
