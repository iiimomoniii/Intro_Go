"# Intro_Go" 
#How to new prepair before run index.go inside Sub-Project
cd Intro
go work init
go work use demo1 or ...

#How to run demo
cd demo..
go run index.go

go mod init demo3
go mod init gin
go mod tidy
go get github.com/gin-gonic/gin
