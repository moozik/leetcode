@echo off
set pName=%1
mkdir "./problems/%pName%"
set file="./problems/%pName%/%pName%.go"

echo package main >> %file%
echo.>> %file%
echo func main() { >> %file%
echo.>> %file%
echo }>> %file%
echo.>> %file%
echo.>> %file%