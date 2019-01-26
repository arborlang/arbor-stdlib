export GO111MODULE=on

all: test plug

plug: 
	go build -o stdlib.so -buildmode=plugin .

test:
	go test -v .

refresh:
	rm -rf vendor/
	go mod vendor

clean_plugin: refresh plug