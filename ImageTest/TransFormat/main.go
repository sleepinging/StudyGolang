package main

import (
	"fmt"
	"os"

	"image/png"
	"golang.org/x/image/bmp"
)

func test() {
	srcfile, dstfile := `E:\temp\1.png`, `E:\temp\1.bmp`
	file, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create(dstfile)
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	//img, err := png.Decode(file) //解码
	img, err := png.Decode(file) //解码
	if err != nil {
		fmt.Println(err)
	}
	bmp.Encode(file1, img)
	//jpeg.Encode(file1, img, &jpeg.Options{100}) //编码，图像质量100
}

func main() {
	test()
}
