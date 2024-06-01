@echo off
set GO_FILE=nft_mint.go

:loop
go run %GO_FILE%
if %ERRORLEVEL% equ 0 (
    echo Program exited successfully. Restarting in 3 seconds...
    timeout /t 3 /nobreak
    goto loop
) else (
    echo Program exited with an error. Not restarting.
)
pause
