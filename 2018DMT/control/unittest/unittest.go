package unittest

import (
	"../../dao"
	"../../global"
	"../../models"
	"../../tools"
	"../EmailVerify"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

func TestEmailVerify() {
	email := "237731947@qq.com"
	//EmailVerify.UpdateCode(email, code)
	code := EmailVerify.SendCode(email)
	fmt.Println(code)
	fmt.Scanln(&code)
	//fmt.Println(code)
	fmt.Println(EmailVerify.CheckCode(email, code))
}

func getroot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), r.RemoteAddr, "访问")
	pt := global.CurrPath
	fs := http.FileServer(http.Dir(pt + `\view\wwwroot\`))
	fs.ServeHTTP(w, r)
}

func webtest() {
	http.HandleFunc("/", getroot)
}

func PostTest() {
	url := "http://193.112.77.180/register/sendcode"
	data := "Email=237731947@qq.com"
	res, err := tools.HttpPost(url, data, nil)
	tools.ShowErr(err)
	fmt.Println(res)
}

func LoginTest() {
	e := dao.CheckLogin(&models.Login{Email: "237731947@qq.com", Password: "123456"})
	fmt.Println(e)
}

func jsonTest() {
	str := `{"email":"23@2","password":"123"}`
	o := &models.Login{}
	err := json.Unmarshal([]byte(str), o)
	tools.ShowErr(err)
	fmt.Println(o)
}

func PostLoginTest() {
	time.Sleep(time.Second * 3)
	url := "http://127.0.0.1/login"
	data := "Email=237731947@qq.com&Password=123456"
	res, err := tools.HttpPost(url, data, nil)
	tools.ShowErr(err)
	fmt.Println(res)
}

func ConfigTest() {
	file, _ := os.OpenFile(global.CurrPath+`config.json`, os.O_RDONLY, os.ModeType)
	defer file.Close()
	bs, _ := ioutil.ReadAll(file)
	cfg := models.Config{}
	json.Unmarshal(bs, &cfg)
	fmt.Println(cfg)
}

func jiamiTest() {
	mi, err := tools.Encrypt("测试", global.MiKey)
	fmt.Println(mi, err)
	str, err := tools.Decrypt(mi, global.MiKey)
	fmt.Println(str, err)
}

func CookieTest() {
	str := dao.GenUserCookie("237731947@qq.com")
	fmt.Println(dao.GetUserIdFromCookie(str))
}

type TT struct {
	A string `gorm:"123" json:"ww"`
	B string `json:b`
	C string `json:c`
}

func TestStruct() {
	var t TT
	str := `{"ww":"1","b":"2","c":"3","d":"4"}`
	err := json.Unmarshal([]byte(str), &t)
	fmt.Println(t, err)
	//select last_insert_rowid();
}

func JobTest() {
	job := &models.Job{
		Name:        "洗碗",
		PublisherId: 1,
		Salary:      3000.0,
		//PublishTime: time.Now(),
		Time:        "每周一到周六",
		Weekend:     -1,
		Pickup:      -1,
		Eat:         1,
		Live:        -1,
		WuXianYiJin: 1,
		Place:       "食堂",
		LimPeople:   10,
		NowPeople:   3,
		Sex:         2,
		Phone:       "13222222222",
		Detail:      "不想洗碗，找几个人帮我洗",
	}
	//id,err:=dao.PublishJob(job)
	//fmt.Println(id,err)
	//t1:=time.Now()
	//fmt.Println(dao.GetJobById(job.Id+78))
	//t2:=time.Now()
	//fmt.Println(t2.Sub(t1))
	//jb,err:=dao.ShowJob(5)
	//fmt.Println(jb,err)

	c := dao.QueryJobCount(job)
	jobs := dao.QueryJob(job, 1, 1)
	fmt.Println(c, jobs)

	//err := dao.UpdataJob(6, job)
	//err := dao.DeleteJob(5)
	//fmt.Println(err)

	////事务
	//wg:=new(sync.WaitGroup)
	//t1:=time.Now()
	//for i:=0;i<10;i++{
	//	wg.Add(1)
	//	go func(){
	//		j:=new(models.Job)
	//		j.CopyJobFromEId(job)
	//		fmt.Println(dao.DeleteJob(i+250))
	//		//fmt.Println(dao.PublishJob(j))
	//		wg.Done()
	//	}()
	//	time.Sleep(time.Millisecond*1)
	//}
	//wg.Wait()
	//t2:=time.Now()
	//fmt.Println(t2.Sub(t1))
}

func reflectTest() {
	t := TT{"1", "2", "3"}
	t2 := TT{"10", "20", "30"}
	s1 := reflect.ValueOf(&t).Elem()
	s2 := reflect.ValueOf(&t2).Elem()
	typ := reflect.TypeOf(t)
	n := typ.NumField()
	for i := 0; i < n; i++ {
		//name := s1.Type().Field(i).Name
		//t := s1.Field(i).Type()
		v := s1.Field(i).Interface()
		switch tp := v.(type) {
		case string:
			s2.Field(i).SetString(v.(string))
		case int:
			s2.Field(i).SetInt(int64(v.(int)))
		default:
			fmt.Println(tp)
		}
		//fmt.Println(name, t, v)
	}
	//// 反射获取测试对象对应的struct枚举类型
	//s := reflect.ValueOf(&t).Elem()
	//// 内置常用类型的设值方法，利用Field序号get
	//s.Field(0).SetString("55")
	fmt.Println(t2)
}

func UserTest() {
	u := &models.User{
		Email:    "123456@qq.com",
		Birthday: time.Now(),
	}
	id, err := dao.AddUser(u)
	fmt.Println(id, err)
	//u, err := dao.GetUserByEmail("237731947@qq.com")
	//fmt.Println(u, err)
	//u2:=&models.User{Id:5}
	//u2.CopyUserFromExpt(u,[]string{"Id",""})
	//fmt.Println(u2)
	u.Name = "麻花疼"
	//u2:=new(models.User)
	//u2.CopyUserFromExpt(u,[]string{"Id"})
	//fmt.Println(dao.UpDateUserInfo(id,u2))
	//tools.ShowErr(dao.UpDateUserInfo(u.Id,u))
	tools.ShowErr(dao.DeleteUser(id))
}

func JobTypeTest() {
	ok, t1, t2, t3 := global.FindJobTypeByName("环保技术")
	fmt.Println(ok, t1.Description, t2.Name, t3.JobName)
}

func ColorTest() {
	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func Test() {
	//ColorTest()
	//UserTest()
	//reflectTest()
	//JobTest()
	//TestStruct()
	//CookieTest()
	//jiamiTest()
	//ConfigTest()
	//go PostLoginTest()
	//jsonTest()
	//LoginTest()
	//PostTest()
	//webtest()
	//TestEmailVerify()
}
