#!/bin/bash

# This script implements starting the hookserver

/hookserver/main --secret $HOOK_SECRET --hook '/bin/bash /scripts/update.sh' --timeout $HOOK_TIMEOUT &
