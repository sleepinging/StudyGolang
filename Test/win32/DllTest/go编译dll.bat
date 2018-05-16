@echo off
set filepath=%1
set dllpath=%~dp1%~n1.dll

IF NOT EXIST "%filepath%" (
	goto err
) ELSE (
	goto make
)

:err
echo 文件不存在！
set/p exepath=请拖拽文件这里或输入文件路径然后按下回车
IF NOT EXIST "%filepath%" (
	goto err
) ELSE (
	goto make
)

:make
echo 编译 %~nx1 -^> %~n1.dll
call go build -ldflags "-s -w" -buildmode=c-shared -o "%dllpath%" "%filepath%"
echo 编译完成!按下回车关闭窗口

:end
pause