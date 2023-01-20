build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip:
	zip sendmessage-main.zip bootstrap

build_zip: build zip

clean:
	rm ./bootstrap
	rm ./sendmessage-main.zip
