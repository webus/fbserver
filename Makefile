# https://github.com/spf13/hugo/blob/master/Makefile
# https://peter.bourgon.org/go-in-production/
all: deps build test
build:
	go build -o fbserver .
test:
	go test .
deps:
	go get .
