FROM golang:1.12

RUN curl https://glide.sh/get | sh

ENV PKG_NAME=github.com/Valkyrie00/curriculum-telegram
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY glide.yaml glide.lock $PKG_PATH/
RUN glide install

COPY . $PKG_PATH
RUN go build && go install

WORKDIR $PKG_PATH