#!/bin/bash

# Root of repo
ROOT_DIR='.'

# Repo full go path
# REPO='github.com/hyperledger/cacti'
REPO='github.com/petermetz/go-monorepo-test'

GOMODULE_PATHS=("pkg/a"
"pkg/b"
"pkg/c")

VERSION=${1:-"2.0.0"}

echo "REPO: $REPO"
echo "VERSION: $VERSION"

MAJOR_VER=""
if [ "${VERSION:0:1}" -gt "1" ]; then
  MAJOR_VER="/v${VERSION:0:1}"
fi

# install go-checksum
go install github.com/vikyd/go-checksum@latest

for GOMODULE in ${GOMODULE_PATHS[@]}; do
  echo "############# START $GOMODULE ################"
  cd $ROOT_DIR/$GOMODULE
  make run-vendor
  GOMOD_DEPS=$((go mod graph | grep "$REPO/$GOMODULE$MAJOR_VER $REPO" | cut -d ' ' -f 2) || (make undo-vendor && echo "ERROR: In generating dependency graph" && exit 1))
  make undo-vendor
  cd - > /dev/null

  for GOMOD_DEP in ${GOMOD_DEPS[@]}; do
    echo "--------- START DEP -----------"
    GOMOD_PATH=$(echo $GOMOD_DEP | awk -F "$MAJOR_VER@" '{print $1}' | awk -F "$REPO/" '{print $2}')
    echo DEP: $GOMOD_DEP
    echo DEP: $GOMOD_PATH
    cp $ROOT_DIR/LICENSE $ROOT_DIR/$GOMOD_PATH
    cd $ROOT_DIR/$GOMOD_PATH
    GOMOD_NAME="$REPO/$GOMOD_PATH$MAJOR_VER"
    if [ ! -f VERSION ]; then
      echo "INFO: VERSION absent"
      cd -
      echo "------------ END --------------"
      continue
    fi
    GOMOD_VERSION=v$(cat VERSION)
    GOMOD_SUM=$(go-checksum . $GOMOD_NAME@$GOMOD_VERSION | grep "GoCheckSum" | cut -d ' ' -f 2 | cut -d '"' -f 2)
    GOMOD_DOTMOD_SUM=$(go-checksum go.mod | grep "GoCheckSum" | cut -d ' ' -f 2 | cut -d '"' -f 2)
    GOMOD_SUM_ENTRY="$GOMOD_NAME $GOMOD_VERSION $GOMOD_SUM"
    GOMOD_DOTMOD_SUM_ENTRY="$GOMOD_NAME $GOMOD_VERSION/go.mod $GOMOD_DOTMOD_SUM"
    echo GOSUM: $GOMOD_SUM_ENTRY
    echo GOSUM: $GOMOD_DOTMOD_SUM_ENTRY
    cd - > /dev/null
    rm $ROOT_DIR/$GOMOD_PATH/LICENSE
    
    cd $ROOT_DIR/$GOMODULE
    UPDATE=false
    (cat go.mod | grep -q "$GOMOD_NAME $GOMOD_VERSION") || UPDATE=True
    if $UPDATE; then
      go mod edit -require $GOMOD_NAME@$GOMOD_VERSION
    else
      echo "ERROR: Version $GOMOD_VERSION already there in go.mod, skipping $GOMOD_PATH in $GOMODULE"
    fi
    UPDATE=false
    (cat go.sum | grep -q "$GOMOD_SUM_ENTRY") || UPDATE=True
    (cat go.sum | grep -q "$GOMOD_DOTMOD_SUM_ENTRY") || UPDATE=True
    if $UPDATE; then
      # mv go.sum go.sum.old
      # grep -v "$GOMOD_NAME $GOMOD_VERSION" go.sum.old > go.sum
      echo $GOMOD_SUM_ENTRY >> go.sum
      echo $GOMOD_DOTMOD_SUM_ENTRY >> go.sum
    else
      echo "ERROR: Version $GOMOD_VERSION already there in go.sum, skipping $GOMOD_PATH in $GOMODULE"
    fi
    cd - > /dev/null
    echo "------------ END --------------"
  done
  echo "############# END $GOMODULE ################\n"
done

