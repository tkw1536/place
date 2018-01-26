FROM nginx
MAINTAINER Tom Wiesing <tom@tkw01536.de>

# Install git and go on top of what we already have
RUN apt-get update && apt-get -y install git-core golang-go rsync

# Add nginx config and scripts
ADD nginx.conf /etc/nginx/conf.d/default.conf
ADD scripts/ /scripts/

# Add the hook server
ADD hookserver/ hookserver/
RUN go build -o /hookserver/main /hookserver/main.go

# environment
ENV DATA_DIR /data
ENV HOOK_TIMEOUT 600
ENV HOOK_SECRET secret


# volumes and endpoint
VOLUME /data
VOLUME /var/www/html
EXPOSE 80

ENTRYPOINT ["/bin/bash", "/scripts/entry.sh"]
