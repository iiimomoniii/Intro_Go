"# Intro_Go" 
#How to prepair before run index.go inside Sub-Project
cd Intro
go work init
go work use demo1 or ...

#How to prepaire before run Sub-Project
go mod init demo3
go mod init gin
go mod tidy
go get github.com/gin-gonic/gin

#How to run isside demo...
cd demo..
go run index.go
