FROM golang:1.13.4 AS SERVERBUILDER

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build .

FROM alpine:latest  
WORKDIR /root/
COPY --from=SERVERBUILDER /go/src/app/sit-o-mat .
CMD ["./sit-o-mat"]