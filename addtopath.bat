@echo off
setlocal

:: Check for administrative privileges
>nul 2>&1 "%SYSTEMROOT%\system32\cacls.exe" "%SYSTEMROOT%\system32\config\system"

if '%errorlevel%' NEQ '0' (
    echo Requesting administrative privileges...
    goto UACPrompt
) else ( goto gotAdmin )

:UACPrompt
    echo Set UAC = CreateObject^("Shell.Application"^) > "%temp%\getadmin.vbs"
    echo UAC.ShellExecute "%~s0", "", "", "runas", 1 >> "%temp%\getadmin.vbs"

    "%temp%\getadmin.vbs"
    exit /B

:gotAdmin
    if exist "%temp%\getadmin.vbs" ( del "%temp%\getadmin.vbs" )
    pushd "%CD%"
    CD /D "%~dp0"

set "sourcePath=%~dp0mapparser.exe"
set "destPath=C:\mapparser\mapparser.exe"

if not exist "C:\mapparser" mkdir "C:\mapparser"

copy /Y "%sourcePath%" "%destPath%"
if errorlevel 1 (
    echo Failed to copy mapparser.exe.
    exit /b 1
)

setx PATH "%PATH%;C:\mapparser" /M
if errorlevel 1 (
    echo Failed to add mapparser to PATH.
    exit /b 1
)

echo mapparser.exe has been successfully copied and added to PATH.
pause