.PHONY: build

build:
	@docker build --no-cache -t email_service:1.1 . 