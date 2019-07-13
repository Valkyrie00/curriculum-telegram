FROM balenalib/raspberrypi3-golang:latest-build AS build

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ENV PKG_NAME=github.com/Valkyrie00/curriculum-telegram
ENV PKG_PATH=$GOPATH/src/$PKG_NAME

COPY Gopkg.toml Gopkg.lock $PKG_PATH/
COPY . $PKG_PATH/

WORKDIR $PKG_PATH

RUN dep ensure
RUN go build && go install

CMD ["curriculum-telegram"]