FROM alpine
MAINTAINER Tom Wiesing <tom@tkw01536.de>

# Install ca-certificates and openssh
RUN apk --no-cache --no-progress add ca-certificates openssh && \
    rm -rf /var/cache/apk/*

ADD cmd/ /place/cmd
ADD place /place/place
ADD server /place/server
ADD updater /place/updater
ADD utils   /place/utils
ADD Makefile /place/Makefile
WORKDIR /place/

# Install make, build, and exit
RUN apk --no-cache --no-progress add --virtual build-deps make go git bash musl-dev && \
    make all && \
    rm -rf $HOME/go && \
    apk --no-progress del build-deps && \
    rm -rf /var/cache/apk/*

# environment
ENV DATA_DIR /data
ENV HOOK_TIMEOUT 600
ENV HOOK_SECRET secret

# volumes and endpoint
VOLUME /data
VOLUME /var/www/html
EXPOSE 80

ENTRYPOINT ["/place/bin/place"]
