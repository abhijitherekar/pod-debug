FROM herekar/abhi-go:1.0

RUN apk update --no-cache \
    && apk upgrade --no-cache \
    && apk add --no-cache \
        libnetfilter_queue \
        iptables \
        ca-certificates \
        iproute2 \
        conntrack-tools \
        ipset \
        tcpdump \
    && rm -rf /var/cache/apk/*

ADD server/server /server
#RUN cd "$GOPATH/src/github.com/server" && \
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o server && cp server / && ls -ltr /

#COPY $GOPATH/src/github.com/server/server .

#RUN ls

ENTRYPOINT ["/server"]