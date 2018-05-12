首先修改本目录下的config.json

json格式，不懂请百度

limit表示小于多少时发出提醒

smtp中配置smtp邮箱信息
不懂也先去百度

sender表示你的smtp邮箱账号
pwd表示smtp邮箱的授权码
host表示smtp服务器地址
to表示提示信息要发送给哪个邮箱

openid目前可以通过抓包方式获取
点开微信嘉兴学院校园卡-校园服务-校园卡余额
过滤IP为111.1.22.131，可以看到
GET /card/queryAcc_queryAccount.shtml?openId=这里就是你的openid&wxArea=10354

为了防止学校的腊鸡服务器爆炸，只能讲这么多了

然后双击运行exe挂着即可，或者可以放到云服务器