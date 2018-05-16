#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include "main.h"

struct MyPoint{//类型和顺序必须和go中的MyPoint对应
	GoString Name;
	long long X;
	long long Y;
};

void dump(char *d, size_t len, const char* file) {
	FILE *fp = nullptr;
	fopen_s(&fp, file, "w+b");
	fwrite(d, sizeof(char), len, fp);
	fclose(fp);
}

void test(){
	printf("Sum(1,2)=%d\n",Sum(1,2));
	
	const char* str="测试字符串\n";
	GoString gstr{str,(ptrdiff_t)strlen(str)};//就是go中的string类型
	Show(gstr);
	//char* cgstr = (char*)malloc(gstr.n + 1);
	//if (!cgstr) { /* handle allocation failure */ }
	//memcpy(cgstr, gstr.p, gstr.n);
	//cgstr[gstr.n] = '\0';
	//dump(cgstr,gstr.n,"1.txt");//dump出来用notepad++打发现是utf-8编码
	//free(cgstr);
	
	MyPoint mp{gstr,0,0};
	SetMyPoint((GoUintptr)&mp);
	printf("%d,%d\n",mp.X,mp.Y);
	
	char* pname=nullptr;
	ToNewGBKCStr(mp.Name,&pname);
	printf("%s\n",pname);
	free(pname);
	pname=nullptr;
	
}
void test2(){
	const char* str="http://www.baidu.com";
	GoString gstr{str,(ptrdiff_t)strlen(str)};
	char* res=nullptr;
	HttpGet(gstr,&res);
	printf("%s\n",res);
	free(res);
	res=nullptr;
}
int main(){
	test2();
	test();
	
	const char* str=R"_TWT_({"name":"来自json的点","x":100,"y":200})_TWT_";//C++11中的raw string,防止转义
	GoString gstr{str,(ptrdiff_t)strlen(str)};
	MyPoint mp;
	JsonToMyPoint(gstr,(GoUintptr)&mp);
	char* pname=nullptr;
	ToNewGBKCStr(mp.Name,&pname);
	printf("%s,%d,%d\n",pname,mp.X,mp.Y);
	free(pname);
	pname=nullptr;
	
	getchar();
	return 0;
}