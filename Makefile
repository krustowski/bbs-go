#
# bbs-go / Makefile
#

include .env.example
-include .env

APP_NAME?=${PROJECT_NAME}

export

all: run

version:
	sed -i "s|^\(PROJECT_VERSION\).*|\1=${PROJECT_VERSION}|" .env.example

run: version
	go build bbs-go
	./bbs-go

push:
	git tag -m "v${PROJECT_VERSION}" "v${PROJECT_VERSION}"
	git push --follow-tags

docker: version
	docker compose up -d --build --remove-orphans --force-recreate
