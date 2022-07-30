FROM golang:1.17-alpine as builder

ENV APP_SRC_DIR=/go/src
ENV GOPATH=/go

RUN mkdir -p $APP_SRC_DIR/app

WORKDIR $APP_SRC_DIR/app

ADD go.mod go.sum $APP_SRC_DIR/

RUN go mod download

ADD . $APP_SRC_DIR/app/

COPY . .

RUN go build -v -o app ./cmd/app && chmod +x app

RUN go build -v -o migrate ./cmd/migrate && chmod +x migrate


FROM alpine:3.13

COPY --from=builder /go/src/app/app /app/
COPY --from=builder /go/src/app/migrate /app/

WORKDIR /app
