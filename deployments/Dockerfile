FROM golang:1.15.2-alpine3.12 AS builder

LABEL maintainer="Vito Castellano <vito.castellano@gmail.com>"

RUN mkdir -p /go/src \
    && mkdir -p /go/bin \
    && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# copy projects
RUN mkdir -p $GOPATH/src/github.com/Valkyrie00/curriculum-telegram
COPY . $GOPATH/src/github.com/Valkyrie00/curriculum-telegram

# go build
WORKDIR $GOPATH/src/github.com/Valkyrie00/curriculum-telegram/cmd
RUN CGO_ENABLED=0 go build -tags=curriculum-telegram -a -installsuffix cgo -o /go/bin/curriculum-telegram .


FROM alpine:3.12

LABEL maintainer="Vito Castellano <vito.castellano@gmail.com>"

RUN apk add --update coreutils tzdata && rm -rf /var/cache/apk/*

WORKDIR /

COPY --from=builder /go/bin/curriculum-telegram /curriculum-telegram

STOPSIGNAL SIGTERM

CMD ["sh", "-c", "/curriculum-telegram"]