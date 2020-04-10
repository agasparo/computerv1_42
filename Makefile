include .env

init:
	go get github.com/fatih/color
	go build -o $(PROJECTNAME)

clean: 
	rm $(PROJECTNAME)

test:
	go get github.com/fatih/color
	go run src/test/test.go

all: init