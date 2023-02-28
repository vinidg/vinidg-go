FROM golang:1.20-alpine as build_go

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/goapp .

FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_go /app/out/goapp /app/goapp

EXPOSE 3000

CMD ["/app/goapp"]