container_push:
	cd server;GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o server;cd ../
	docker build -t gcr.io/aporetodev/abhi-server:11 -f docker-server/Dockerfile .
	docker push gcr.io/aporetodev/abhi-server:11