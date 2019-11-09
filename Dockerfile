FROM golang:1.13.4 AS SERVERBUILDER

WORKDIR /go/src/app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o sit-o-mat


FROM scratch
WORKDIR /go/bin/
COPY --from=SERVERBUILDER /go/src/app/sit-o-mat .
COPY --from=SERVERBUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/sit-o-mat"]