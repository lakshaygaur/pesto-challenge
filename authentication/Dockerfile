FROM golang:1.22.2-alpine as builder

WORKDIR /home/build

COPY . .

RUN go mod tidy
RUN apk add make

RUN make run-build

FROM alpine 
#go:1.22.0-alpine

WORKDIR /home/app

COPY --from=builder /home/build/build/* .

CMD ["/home/app/authentication"]