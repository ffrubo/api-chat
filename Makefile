STACK_NAME := chat

build:
	cd connect && $(MAKE) build_zip
	cd sendmessage && $(MAKE) build_zip
	cd serversendmessage && $(MAKE) build_zip

clean:
	cd connect && $(MAKE) clean
	cd sendmessage && $(MAKE) clean
	cd serversendmessage && $(MAKE) clean

create-stack:
	aws cloudformation create-stack \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_IAM \
		--template-body file://template.yaml

update-stack:
	aws cloudformation update-stack \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_IAM \
		--template-body file://template.yaml
