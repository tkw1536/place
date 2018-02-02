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

# listen on port 80
EXPOSE 80
ENV BIND_ADDRESS 0.0.0.0:80

# store the ssh key in /data/id_rsa
VOLUME /data/
ENV SSH_KEY_PATH /data/id_rsa

# server the webhook under /webhook/
ENV WEBHOOK_PATH /webhook/

# and setup static hosting under /var/www/html
VOLUME /var/www/html
ENV STATIC_PATH /var/www/html

ENV GIT_BRANCH master
ENV GIT_CLONE_TIMEOUT 600


# to be provided by user:
# ENV GIT_URL git@github.com:example/domain.tld.git
# ENV GITHUB_SECRET supersecret
# ENV GITLAB_SECRET supersecret
# ENV DEBUG 1

ENTRYPOINT ["/place/bin/place"]
