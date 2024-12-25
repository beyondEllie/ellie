.PHONY: build install

# Building the ellie binary
build:
	go build -o ellie

# Installing the ellie binary to /usr/local/bin
install: build
	sudo cp ellie /usr/local/bin/ellie
