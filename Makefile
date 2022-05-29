.PHONY: build test

build:
	go build -o ./build/

test:
	go test ./...

remove-app-config:
	rm -f application.yaml

copy-stg-config: remove-app-config
	cp ./config/yamls/application.stg.yaml application.yaml

copy-prod-config: remove-app-config
	cp ./config/yamls/application.prod.yaml application.yaml

build-stg: copy-stg-config build

build-prod: copy-prod-config build
