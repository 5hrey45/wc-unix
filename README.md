# Unix wc tool written in Go

# How to use?
Use go build \
run the executable produces and provide the filename as argument

E.G. go build -o ccwc.exe \
ccwc.exe -flags (if any) filename

Alternatively you can use go run \
go run main.go -flags (if any) filename

You can also use flags available in wc tool such as
- -c for bytes
- -m for characters
- -w for words
- -l for lines