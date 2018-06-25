@echo off

Set pn=dmt2018_t.exe

echo 检测进程%pn%中

:LOOP
tasklist /nh|find /i "%pn%">nul
if ERRORLEVEL 1 (
    echo %date:~0,10% %time:~0,8%
    echo 进程挂了，重启一下
    call "%pn%"
) else (
    Set temp=1
)
ping 127.0.0.1 -n 6 >nul
goto LOOP

:NOPROCESS


pause