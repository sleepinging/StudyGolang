var system = require('system');
var page = require('webpage').create();

//page.customHeaders={"User-Agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36"};
//var url="http://www.wanfangdata.com.cn/search/downLoad.do?page_cnt=13&language=&resourceType=patent&source=WF&resourceId=CN201710644621.7&resourceTitle=%E7%9D%A1%E8%A7%89%E5%8F%AB%E9%86%92%E7%B3%BB%E7%BB%9F&isoa=&type=patent";
var url=system.args[1];
page.open(url, function(status) {
	//console.log(status);
	if ( status === "success" ) {
		//page.render('1.png');
		var content = page.evaluate(function () {
			var dlurl=document.getElementById("downloadIframe").getAttribute("src");
			//console.log(dlurl);
			return dlurl;
		});
		console.log(content)
		
	}
    phantom.exit(0);
});