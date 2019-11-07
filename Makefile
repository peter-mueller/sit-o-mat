build: 
	go build -o bin/sit-o-mat main.go
	go build -o bin/sitomatadmin admincli/sitomatadmin/main.go
	cp admincli/sitomatadmin/environment.sh bin
	cp admincli/sitomatadmin/example.sh bin

init: build 
	cd bin && ./example.sh

assign: 
	curl http://admin:password@localhost:8080/sitomat 

run: build stop
	./bin/sit-o-mat &

stop:
	pkill sit-o-mat || true