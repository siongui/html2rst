export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


default:
	@echo "\033[92mIterating HTML and convert HTML to rst ...\033[0m"
	@go test -v

example:
	@echo "\033[92mRunning example ...\033[0m"
	@cd usage; go run example.go

fmt:
	@echo "\033[92mGo fmt source code ...\033[0m"
	@go fmt *.go
	@go fmt usage/*.go

install:
	@echo "\033[92mInstalling golang.org/x/net/html ...\033[0m"
	go get -u golang.org/x/net/html
	@echo "\033[92mInstalling golang.org/x/text/width ...\033[0m"
	go get -u golang.org/x/text/width
	@echo "\033[92mInstalling github.com/djimenez/iconv-go ...\033[0m"
	go get -u github.com/djimenez/iconv-go

self:
	@echo "\033[92mInstalling github.com/siongui/html2rst ...\033[0m"
	go get -u github.com/siongui/html2rst
