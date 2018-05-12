package main

import (
	"twt/nettools"
	"fmt"
)

func main() {
	headers := map[string]string{};
	headers["referer"] = "https://servicewechat.com/wxee55405953922c86/96/page-frame.html"
	headers["Content-Type"] = "application/json"
	res, err := nettools.HttpPost("https://api-xcx-qunsou.weiyoubot.cn/xcx/checkin/v3/doit",
		`{"cid":"5aea7b7be17b4a3574115d38","formid":"1525996436479","text":"签到","pic":[],"audio":[],"audio_len":[],"longitude":null,"latitude":null,"access_token":"e6b055a5edd44c26a8b3424189425466"}`,
		headers)
	if err != nil {
		return
	}
	fmt.Println(res)
}
