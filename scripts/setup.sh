#!/bin/bash

# This script implements the container setup procedure and is called upon entry
# into the container

# create ssh key, if it doesn't already exist
if [ ! -f "$SSH_ID_FILE" ]; then
    echo "$SSH_ID_FILE does not exist, generating a new one"
    echo "============================================"
    ssh-keygen -t rsa -N '' -f $SSH_ID_FILE
    echo "============================================"
fi

# setup ssh config
mkdir -p $HOME/.ssh
echo "IdentityFile $SSH_ID_PUB_FILE" > $HOME/.ssh/config
echo "StrictHostKeyChecking=no" >> $HOME/.ssh/config

# fix permissions
chmod 600 $SSH_ID_FILE
chmod 600 $SSH_ID_PUB_FILE

# print out the signature
echo "Your ssh key signature is: "
echo "============================================"
cat $SSH_ID_PUB_FILE
echo "============================================"
echo "Add this to your repository for ssh access. "


# and start an ssh agent
ssh-agent > $SSH_AGENT_FILE
source $SSH_AGENT_FILE
ssh-add $SSH_ID_FILE
