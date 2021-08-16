#!/usr/bin/env bash

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

mkdir -p output/$PathPre/bin
cp script/* output/
chmod +x output/bootstrap.sh

cd ./$PathPre
pwd

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -o $CURDIR/output/$PathPre/bin/${RUN_NAME}
else
    go test -c -covermode=set -o $CURDIR/$PathPre/output/bin/${RUN_NAME} -coverpkg=./...
fi