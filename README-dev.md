# Dev Readme

## Commands
```bash
# init go module
go mod init github.com/adegoodyer/ssh-load-tester && \
go mod tidy

# get packages
go get golang.org/x/crypto/ssh

# local install
go install ssh-load-tester.go

# run
ssh-load-tester
```
