package main
import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"fmt"
)
func main() {
	mainWindow.Run()
}
var LableHello=Label{
	Text: "欢迎使用!",
}

var inname, inpwd *walk.TextEdit
var inputname=TextEdit{
	AssignTo: &inname,
}
var inputpwd=TextEdit{
	AssignTo: &inpwd,
}
var inputnamewidget=[]Widget{
	labelname,
	inputname,
}
var labelname=Label{
	Text:"用户名",
}
var inputpwdwidget=[]Widget{
	labelpwd,
	inputpwd,
}
var labelpwd=Label{
	Text:"密码",
}
var widget=[]Widget{
	LableHello,
	HSplitter{
		Children:inputnamewidget,
	},
	HSplitter{
		Children:inputpwdwidget,
	},
	PushButton{
		Text:"登录",
		OnClicked:OnLogin,
	},
}
var mainWindow=MainWindow{
	Title:"登录程序",
	//MinSize:Size{400, 200},
	Size:Size{400,100},
	Layout:VBox{},
	Children:widget,
}

func OnLogin(){
	fmt.Println(inname.Text())
	fmt.Println(inpwd.Text())
	inpwd.SetText("*****")
}