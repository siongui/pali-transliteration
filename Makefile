# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH

GO_VERSION=1.12.17
ifndef TRAVIS
	# set environment variables on local machine or GitLab CI
	export GOROOT=$(realpath ../paligo/go)
	export GOPATH=$(realpath ../paligo/)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

devserver: fmt js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver/devserver.go

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build gopherjs/*.go -m -o gopherjs/app.js

test: fmt
	@go test -v

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt gopherjs/*.go

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@wget https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz
	@tar -xvzf go$(GO_VERSION).linux-amd64.tar.gz
	@rm go$(GO_VERSION).linux-amd64.tar.gz

install:
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/siongui/godom
	go get -u github.com/siongui/pali-transliteration
	go get -u github.com/siongui/gopalilib/libfrontend/velthuis
