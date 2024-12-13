#
# bbs-go / Dockerfile
#

#
# stage 1 --- build go binary
#

# https://hub.docker.com/_/golang

ARG GOLANG_VERSION "1.20"
FROM golang:${GOLANG_VERSION}-alpine AS bbs-staging

ARG APP_NAME

ENV APP_NAME ${APP_NAME}
RUN apk --no-cache add tzdata git

WORKDIR /go/src/${APP_NAME}
COPY . .

RUN go mod tidy
RUN go build ${APP_NAME}


#
# stage 2 --- run single binary in raw alpine
#

FROM alpine:3.17 AS bbs-runtime

ARG LISTEN_ADDR
ARG LISTEN_PORT
ARG SWAPI_TOKEN
ARG TZ

ENV LISTEN_ADDR ${LISTEN_ADDR}
ENV LISTEN_PORT ${LISTEN_PORT}
ENV SWAPI_TOKEN ${SWAPI_TOKEN}
ENV TZ ${TZ}

COPY --from=bbs-staging /go/src/bbs-go/bbs-go /usr/bin/bbs-go

USER nobody
EXPOSE ${LISTEN_PORT}
CMD bbs-go

