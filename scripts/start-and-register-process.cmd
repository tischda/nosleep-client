:: ----------------------------------------------------------------------------
:: Start and register process
:: ----------------------------------------------------------------------------
@echo off
setlocal

set PORT=9001
set ID=%random%_%~nx0

echo -- get process id from this script
title %ID%
for /f "tokens=2 delims=," %%i in ('tasklist /v /fo csv ^| findstr /i "%ID%"') do set "CMDPID=%%~i"
echo Current batch PID: %CMDPID%
echo.

echo -- check if server is running
nosleep-client.exe --port %PORT% Read 2>nul
if errorlevel 1 (
	echo Client cannot not connect.
	echo.
	echo -- start nosleep-server
	start /b "nosleep-server" nosleep-server.exe --port %PORT%
	ping -n 2 127.0.0.1 >nul
)

echo.
echo -- register process %CMDPID%
nosleep-client.exe --port %PORT% Register --pid %CMDPID%
echo.

echo -- read current status
nosleep-client.exe --port %PORT% Read
echo.

REM simulate long running process...
ping -n 5 127.0.0.1 >nul

echo -- unregister process %CMDPID%
nosleep-client.exe --port %PORT% Unregister --pid %CMDPID%
