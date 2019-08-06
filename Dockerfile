FROM golang:1.12

ENV go111module=on

WORKDIR /tmp
RUN go get github.com/izumin5210/gex/cmd/gex
RUN curl https://github.com/izumin5210/gex/releases/download/v0.5.0/gex_linux_amd64.tar.gz -L -o gex_linux_amd64.tar.gz
RUN tar zxvf gex_linux_amd64.tar.gz
RUN cp gex_linux_amd64/gex /usr/local/bin

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN gex grapi build
CMD /app/bin/server
