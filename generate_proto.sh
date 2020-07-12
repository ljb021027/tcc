#!/usr/bin/env bash
#set -x
set -e
cur_script_dir="`cd $(dirname $0) && pwd`"
WORK_HOME="${cur_script_dir}/proto"
IMPORT_HOME="$GOPATH/src"
echo "dirname $WORK_HOME"
echo "IMPORT_HOME: $IMPORT_HOME"
echo "WORK_HOME = $WORK_HOME"
find $WORK_HOME -name "*.proto" | while read proto; do
  dir="`dirname $proto`"
  echo "dir: `cd $dir && pwd`"
  docker run --rm -v $dir:/defs -v ${IMPORT_HOME}:/input -v $WORK_HOME:/workspace blademainer/protoc-all:1.23_v0.0.3 -i /defs -i /input -d /defs/ -l go -o /workspace/go --validate-out "lang=go:/workspace/go" --with-gateway --lint $addition;
done