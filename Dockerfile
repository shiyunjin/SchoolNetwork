FROM golang:1.11.2 as go-builder

# work in gopath dir
WORKDIR /go/src/github.com/shiyunjin/Labs-Gate/

ADD ./ .

ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN /usr/local/go/bin/go get ./...
RUN /usr/local/go/bin/go build -o lab-gate .

FROM node:8.12 as react-builder

WORKDIR /app

RUN mkdir -p /app/view && \
    cd /app/view && \
    git clone https://github.com/shiyunjin/Labs-Gate-UI.git . && \
    npm install && \
    npm run build

FROM scratch

ENV APP_DIR=/app

ADD config.tson /app/config.tson
ADD run.sh      /app/run.sh

COPY --from=go-builder /go/src/github.com/shiyunjin/Labs-Gate/lab-gate /app/lab-gate
COPY --from=react-builder /app/view/build /app/system/view/build

WORKDIR $APP_DIR

EXPOSE 8080

CMD ["/app/run.sh"]
