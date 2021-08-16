#! /usr/bin/env bash
export PSM=${PSM:-kitex.service.c}
CURDIR=$(cd $(dirname $0); pwd)

PathPre=

case $RUN_ENV in
server_a)
  PathPre="KiteX-A"
  RUN_NAME="kitex.service.a"
  ;;

server_b)
  PathPre="KiteX-B"
  RUN_NAME="kitex.service.b"
  ;;

server_c)
  PathPre="KiteX-C"
  RUN_NAME="kitex.service.c"
  ;;

*)
  echo "缺少环境变量RUN_ENV "
  exit 1
  ;;
esac

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

RUNTIME_ROOT=$RUNTIME_ROOT/$PathPre

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

echo "$CURDIR/$PathPre/bin/$RUN_NAME"

exec "$CURDIR/$PathPre/bin/$RUN_NAME"