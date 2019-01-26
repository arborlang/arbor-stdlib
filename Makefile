all: test plug

plug: 
	go build -o stdlib.so -buildmode=plugin .

test:
	go test -v .