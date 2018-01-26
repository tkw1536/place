#!/bin/bash

# This script implements the post-commit hook and is called upon receiving a
# proper post commit hook. It has a timeout specified by $HOOK_TIMEOUT, and will
# automatically be killed after this period.

# cleanup previous clones if they exist
if [ -d "$TEMP_CLONE" ]; then
  rm -rf $TEMP_CLONE
fi

# load ssh agent stuff
source $SSH_AGENT_FILE

# clone the repository
echo git clone --depth 1 $REPO_URL $TEMP_CLONE
git clone --depth 1 $REPO_URL $TEMP_CLONE

# copy everything over to the target directory /var/www
# except for .git
echo rsync -rv --delete-excluded --exclude .git .gitignore $TEMP_CLONE/ $PUBLIC_DIR
rsync -rv --delete-excluded --exclude .git .gitignore $TEMP_CLONE/ $PUBLIC_DIR

# and kill the old clone
rm -rf $TEMP_CLONE
