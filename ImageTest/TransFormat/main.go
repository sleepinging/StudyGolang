package main

import (
	"image/png"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
	"reflect"
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

	img, err := png.Decode(file) //解码
	rec := img.Bounds()
	if err != nil {
		fmt.Println(err)
	}
	//bmphead:=make([]byte,54)
	w, h := rec.Max.X, rec.Max.Y
	linebyte := (((w * 24) + 31) >> 5) << 2
	bmpdata := make([]byte, h*linebyte)
	for y := rec.Min.Y; y < rec.Max.Y; y++ {
		for x := rec.Min.X; x < rec.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			bmpdata[y*linebyte+x*3+0] = uint8(b)
			bmpdata[y*linebyte+x*3+1] = uint8(g)
			bmpdata[y*linebyte+x*3+2] = uint8(r)
		}
	}

	ioutil.WriteFile(dstfile, bmpdata, os.ModeAppend)
	//jpeg.Encode(file1, img, &jpeg.Options{100}) //编码，图像质量100
}

type aaa struct {
	c byte
	a uint16
	b uint64
	d byte
}

func main() {
	a := aaa{0x01, 0x0203, 0x04, 5}
	fmt.Println(unsafe.Sizeof(a))
	typ := reflect.TypeOf(a)
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	dataBytes := (*[12]byte)(unsafe.Pointer(&a))
	fmt.Printf("Bytes are %#v\n", dataBytes)
}
