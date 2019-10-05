# pod-yaml-helpers


## Built the binary image:
cd server;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o server;cd ../
docker build -t gcr.io/aporetodev/abhi-server:tag -f docker-server/Dockerfile .
docker push gcr.io/aporetodev/abhi-server:tag
