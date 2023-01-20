build:
	cd connect && $(MAKE) build_zip
	cd sendmessage && $(MAKE) build_zip
	cd serversendmessage && $(MAKE) build_zip

clean:
	cd connect && $(MAKE) clean
	cd sendmessage && $(MAKE) clean
	cd serversendmessage && $(MAKE) clean
