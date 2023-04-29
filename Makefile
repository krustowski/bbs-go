#
# bbs-go / Makefile
#

include .env.example
-include .env

export

all: run

version:
	sed -i "s|^\(PROJECT_VERSION\).*|\1=${PROJECT_VERSION}|" .env.example

run: version
	go build bbs-go
	./bbs-go

push:
	git tag -m "${PROJECT_VERSION}" "${PROJECT_VERSION}"
	git push --follow-tags
