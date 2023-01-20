build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip:
	zip serversendmessage-main.zip bootstrap

build_zip: build zip

clean:
	rm ./bootstrap
	rm ./serversendmessage-main.zip
