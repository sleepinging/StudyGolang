package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type XmlPoints struct {
	XMLName   xml.Name      `xml:"array"`
	PointsArr []PointsArray `xml:"array"`
}

type PointsArray struct {
	XMLName xml.Name `xml:"array"`
	Points  []string `xml:"string"`
}

func transp(pstr string) (cpstr string) {
	rs := []rune(pstr)
	cpstr = string(rs[1 : len(rs)-1])
	cpstr = `(` + cpstr + `)`
	return
}

func gop2c(filename, vertname, bodyname, cpanme, faname, fcname string) (cppstr string, err error) {
	//vertname, bodyname := "vert", "body"
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	xps := new(XmlPoints)
	err = xml.Unmarshal(bs, xps)
	if err != nil {
		return
	}
	for i, ps := range xps.PointsArr {
		psname := cpanme + " " + vertname + fmt.Sprintf("%d", i) + "[]"
		psname += `={`
		for _, p := range ps.Points[:len(ps.Points)-1] {
			psname += cpanme + transp(p) + `,`
		}
		psname += "Point" + transp(ps.Points[len(ps.Points)-1]) + `};`
		cppstr += psname
		//fmt.Println(psname)
		stname := bodyname + `->` + faname + `(` + fcname + `(` + vertname + fmt.Sprintf("%d", i) +
			`,` + fmt.Sprintf("%d", len(ps.Points)) + `));`
		cppstr += "\n" + stname + "\n"
		//fmt.Println(stname)
	}
	return
}

func main() {
	vertname := flag.String("v", "vert", "每一组多边形的基本变量名")
	bodyname := flag.String("b", "body", "添加到的PhysicsBody变量名")
	filename := flag.String("f", "1.xml", "要解析的文件")
	savename := flag.String("s", "gop2c.cpp", "要保存的文件")
	cpanme := flag.String("cpn", "Point", "Point类名")
	faname := flag.String("fan", "addShape", "addShape方法名")
	fcname := flag.String("fcn", "PhysicsShapePolygon::create", "创建多边形方法名")
	flag.Parse()
	ts := time.Now()
	cppstr, err := gop2c(*filename, *vertname, *bodyname, *cpanme, *faname, *fcname)
	if err != nil {
		panic(err)
	}
	t1 := time.Now()
	err = ioutil.WriteFile(*savename, []byte(cppstr), os.ModeType)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	fmt.Println("读取处理时间:", t1.Sub(ts), "\n保存时间:", t2.Sub(t1), "\n总时间:", t2.Sub(ts))
}
