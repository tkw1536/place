#!/bin/bash

# This script implements the entry point for the hook server
# It is called without any arguments and is reponsible for
# setting up the container, starting the hookserver, and
# starting nginx

export WWW_HOME="/var/www"

export SSH_AGENT_FILE="$WWW_HOME/.ssh/agent"
export SSH_ID_FILE=$DATA_DIR/id_rsa
export SSH_ID_PUB_FILE=$DATA_DIR/id_rsa.pub

export TEMP_CLONE=$DATA_DIR/clone
export PUBLIC_DIR="$WWW_HOME/html"

# chown all the relevant directories
chown -R www-data:www-data $WWW_HOME
chown -R www-data:www-data $DATA_DIR

# Run the setup script
su www-data -s /bin/bash /scripts/setup.sh

# start the hookserver
su www-data -s /bin/bash /scripts/hookserver.sh &

# and start nginx
nginx -g 'daemon off;'
