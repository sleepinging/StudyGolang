package BMP

import (
	"image"
	"image/color"
	"os"
	"unsafe"
	"fmt"
)

//bmp文件头
type BITMAPFILEHEADER struct {
	//位图类别，根据不同的操作系统而不同，
	// 在Windows中，此字段的值总为BM
	BfType uint16

	//BMP图像文件的大小
	BfSize uint32

	//总为0
	BfReserved1 uint16

	//总为0
	BfReserved2 uint16

	//BMP图像数据的地址
	BfOffBits uint32
}

//位图信息头
type BITMAPINFOHEADER struct {
	//本结构的大小，根据不同的操作系统而不同，
	// 在Windows中，此字段的值总为28h字节=40字节
	BiSize uint32

	//BMP图像的宽度，单位像素
	BiWidth uint32

	//总为0
	BiHeight uint32

	//总为0
	BiPlanes uint16

	//BMP图像的色深，即一个像素用多少位表示，
	// 常见有1、4、8、16、24和32，
	// 分别对应单色、16色、256色、16位高彩色、
	// 24位真彩色和32位增强型真彩色
	BibitCount uint16

	//压缩方式，0表示不压缩，1表示RLE8压缩，
	// 2表示RLE4压缩，3表示每个像素值由指定的掩码决定
	BiCompression uint32

	//BMP图像数据大小，必须是4的倍数，
	// 图像数据大小不是4的倍数时用0填充补足
	BiSizeImage uint32

	//水平分辨率，单位像素/m
	BiXPelsPerMeter uint32

	//垂直分辨率，单位像素/m
	BiYPelsPerMeter uint32

	//BMP图像使用的颜色，0表示使用全部颜色，
	// 对于256色位图来说，此值为100h=256
	BiClrUsed uint32

	//重要的颜色数，此值为0时所有颜色都重要，
	// 对于使用调色板的BMP图像来说，
	// 当显卡不能够显示所有颜色时，此值将辅助驱动程序显示颜色
	BiClrImportant uint32
}

//调色板
type RGBQUAD struct {
	//蓝色值
	RGBBlue byte

	//绿色值
	RGBGreen byte

	//红色值
	RGBRed byte

	//保留，总为0
	RGBReserved byte
}

//type BITMAPDATA struct {
//	Data []byte
//	//Data *[]byte
//}

type BITMAP struct {
	Fileheader BITMAPFILEHEADER
	Infoheader BITMAPINFOHEADER
	Rgbq       RGBQUAD
	Data       []byte
}

//将任意颜色转变成他自己的颜色
func (bmp *BITMAP) ColorModel() (cm color.Model) {
	return
}

//图片中非０color的区域
func (bmp *BITMAP) Bounds() (rec image.Rectangle) {
	return
}

//返回指定点的像素color.Color
func (bmp *BITMAP) At(x, y int) (c color.Color) {
	return
}

//从数据生成BMP
func (bmp *BITMAP) GenBitMapImgByData(data *[]byte, width, height uint32, colorbit uint16) {

}

//保存图片
func (bmp *BITMAP) Save(filename string) (err error) {
	var fl *os.File
	fl, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fl.Close()
	n := 0
	n, err = fl.Write(*(*[]byte)(unsafe.Pointer(&bmp.Fileheader)))
	n, err = fl.Write(*(*[]byte)(unsafe.Pointer(&bmp.Infoheader)))
	n, err = fl.Write(*(*[]byte)(unsafe.Pointer(&bmp.Rgbq)))
	n, err = fl.Write(bmp.Data)
	if err != nil {
		return
	}
	if uint32(n) < uint32(58)+bmp.Fileheader.BfSize {
		err = fmt.Errorf("只保存一部分")
	}
	return
}
