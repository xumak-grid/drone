#!/bin/sh
echo "Running exposer"
if EXPORTED_ENVS=$(/bin/exposer)
then
  eval $EXPORTED_ENVS
  export DRONE_HOST="http://$DRONE_HOST"
  export DRONE_GOGS_URL="http://$DRONE_GOGS_URL"
  printenv
  /drone server
else
  exit 1
fi
