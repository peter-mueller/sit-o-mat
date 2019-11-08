FROM golang:1.13.4 AS SERVERBUILDER

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o sit-o-mat

FROM scratch
WORKDIR /go/bin/
COPY --from=SERVERBUILDER /go/src/app/sit-o-mat .
ENTRYPOINT ["/go/bin/sit-o-mat"]