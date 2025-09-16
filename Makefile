# Makefile for Ellie CLI (Go + Rust FFI)

.PHONY: all rust go clean build_linux build_windows build_mac build_android build_linux_upx build_windows_upx build_mac_upx build_android_upx

all: rust go

rust:
	cd rustmods/elliecore && cargo build

go:
	go mod tidy
	CGO_ENABLED=1 go build -o ellie

build_linux:
	@echo 'building linux binary...'
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_linux_amd64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie     

build_linux_upx:
	@echo 'building linux binary with UPX...'
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_linux_amd64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_windows:
	@echo 'building windows executable...'
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ellie_windows_amd64.exe
	@echo 'zipping build...'
	zip binaries/ellie_windows_amd64.zip ellie_windows_amd64.exe
	@echo 'cleaning up...'
	rm ellie_windows_amd64.exe

build_windows_upx:
	@echo 'building windows executable with UPX...'
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ellie_windows_amd64.exe
	@echo 'compressing with UPX...'
	upx --best ellie_windows_amd64.exe
	@echo 'zipping build...'
	zip binaries/ellie_windows_amd64_upx.zip ellie_windows_amd64.exe
	@echo 'cleaning up...'
	rm ellie_windows_amd64.exe

build_mac:
	@echo 'building mac binary...'
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_mac_amd64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_mac_upx:
	@echo 'building mac binary with UPX...'
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_mac_amd64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_android:
	@echo 'building android binary'
	env GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_android_arm64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_android_upx:
	@echo 'building android binary with UPX...'
	env GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_android_arm64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

clean:
	cd rustmods/elliecore && cargo clean
	rm -f ellie
