# Unix wc tool written in Go

## How to use?

### Use go build

E.G. go build -o ccwc.exe \
ccwc.exe -flags (if any) filename

### Alternatively you can use go run
go run main.go -flags (if any) filename

run the .go file or the executable produced by building and (provide the filename as argument or provide the text data through stdin through piping)

cat filename | go run main.go
or
cat filename | ccwc.exe

### You can also use flags available in wc tool such as
- -c for bytes
- -m for characters
- -w for words
- -l for lines