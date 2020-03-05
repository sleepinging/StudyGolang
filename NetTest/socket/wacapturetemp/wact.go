package main

import (
	"net"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		buffer := make([]byte, 2048)
		n,err:=c.Read(buffer)
		if err != nil {
			return
		}
		fmt.Println(n,buffer[:n])
	}
}


func main() {
	l, err :=net.Listen("unix","/tmp/pbwaserver.sock")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}

}
