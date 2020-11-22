#################
# Builder
#################
FROM arm64v8/golang:1.15-alpine AS builder

LABEL maintainer="Vito Castellano <vito.castellano@gmail.com>"

# OS
RUN mkdir -p /go/src \
    && mkdir -p /go/bin \
    && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# ENV
ENV CGO_ENABLED=0
ENV GOARCH=arm64
ENV GOOS=linux

# Project setup
ARG PROJECT_PATH=$GOPATH/src/github.com/Valkyrie00/curriculum-telegram

# copy projects
RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH

# go build
WORKDIR $PROJECT_PATH/cmd
RUN go build -tags=curriculum-telegram -a -installsuffix cgo -o ../curriculum-telegram .


#################
# System
#################
FROM arm64v8/alpine:latest

LABEL maintainer="Vito Castellano <vito.castellano@gmail.com>"

RUN apk add --update coreutils tzdata && rm -rf /var/cache/apk/*

COPY --from=builder /go/src/github.com/Valkyrie00/curriculum-telegram /curriculum-telegram

WORKDIR /curriculum-telegram

STOPSIGNAL SIGTERM

CMD ["sh", "-c", "./curriculum-telegram"]
