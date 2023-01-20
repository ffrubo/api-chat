build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip:
	zip connect-main.zip bootstrap

build_zip: build zip

clean:
	rm ./bootstrap
	rm ./connect-main.zip
