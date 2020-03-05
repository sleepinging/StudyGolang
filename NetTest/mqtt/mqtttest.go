package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"os"
	"strings"
	"crypto/des"
	"crypto/cipher"
)

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func decmsg(data []byte)(res []byte,err error){
	res,err=DesDecrypt(data,[]byte("dx$@shdx"))
	if err != nil {
		return
	}
	fmt.Println(res)
	return
}

func main() {
	fmt.Println("start")
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://27.115.88.42:61883")
	opts.SetClientID("ujgfj")
	opts.SetUsername("wangan")
	opts.SetPassword("etonesystem")
	opts.SetCleanSession(false)

	choke := make(chan [2][]byte, 1024)

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		choke <- [2][]byte{[]byte(msg.Topic()), msg.Payload()}
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("wangan/+", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		incoming := <-choke
		topic:=string(incoming[0])
		if !strings.HasPrefix(topic,"wangan/virtualidentity"){
			continue
		}
		//fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		fmt.Printf("RECEIVED TOPIC: %s \n", topic)
		decmsg([]byte(incoming[1]))
	}
}


