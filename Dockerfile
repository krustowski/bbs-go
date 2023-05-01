#
# bbs-go / Dockerfile
#

# https://hub.docker.com/_/golang

ARG GOLANG_VERSION "1.20"
FROM golang:${GOLANG_VERSION}-alpine

ARG APP_NAME
ARG LISTEN_ADDR
ARG LISTEN_PORT
ARG TZ

ENV APP_NAME ${APP_NAME}
ENV LISTEN_ADDR ${LISTEN_ADDR}
ENV LISTEN_PORT ${LISTEN_PORT}
ENV TZ ${TZ}

RUN apk --no-cache add tzdata git

WORKDIR /go/src/${APP_NAME}
COPY . .

#RUN go mod init ${APP_NAME}
RUN go mod tidy

# build the server
RUN go build ${APP_NAME}

EXPOSE ${LISTEN_PORT}
CMD ./${APP_NAME}

