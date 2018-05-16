#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include <windows.h>
#include "main.h"

struct MyPoint{
	GoString Name;
	int X;
	int Y;
};

void dump(char *d, size_t len, const char* file) {
	FILE *fp = nullptr;
	fopen_s(&fp, file, "w+b");
	fwrite(d, sizeof(char), len, fp);
	fclose(fp);
}

int main(){
	printf("Sum(1,2)=%d\n",Sum(1,2));
	
	const char* str="测试字符串\n";
	GoString gstr{str,(ptrdiff_t)strlen(str)};
	Show(gstr);
	//char* cgstr = (char*)malloc(gstr.n + 1);
	//if (!cgstr) { /* handle allocation failure */ }
	//memcpy(cgstr, gstr.p, gstr.n);
	//cgstr[gstr.n] = '\0';
	//dump(cgstr,gstr.n,"1.txt");
	//free(cgstr);
	
	MyPoint mp;
	SetMyPoint((GoUintptr)&mp);
	
	char* pname=nullptr;
	ToNewCStr(mp.Name,&pname);
	printf("%s\n",pname);
	free(pname);
	pname=nullptr;
	
	printf("%d,%d\n",mp.X,mp.Y);
	
	getchar();
	return 0;
}