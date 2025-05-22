@echo off
cls
taskkill /f /im go.exe >nul 2>&1
go run ./cmd