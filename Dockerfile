FROM golang:1-alpine

# Build dependencies
RUN apk --no-cache --no-progress add ca-certificates openssh go git bash musl-dev \
    && rm -rf /var/cache/apk/*

ADD . /go/src/github.com/tkw1536/place
WORKDIR /go/src/github.com/tkw1536/place
RUN go get -v .
RUN CGO_ENABLED=0 GOOS=linux go build -a github.com/tkw1536/place/cmd/place


FROM alpine:latest  
RUN apk --no-cache add ca-certificates openssl

# Install make, build, and exit
RUN apk --no-cache --no-progress add ca-certificates openssh \
    && rm -rf /var/cache/apk/*

COPY --from=0 /go/src/github.com/tkw1536/place/place  /root/

# listen on port 80
EXPOSE 80
VOLUME /var/www/html

# to be provided by user:
# ENV GIT_URL git@github.com:example/domain.tld.git
# ENV GITHUB_SECRET supersecret
# ENV GITLAB_SECRET supersecret
# ENV DEBUG 1

ENTRYPOINT ["/root/place"]
